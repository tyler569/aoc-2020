#!/usr/bin/env ruby

require 'optparse'

options = {
  input: "test_input",
}

OptionParser.new do |opt|
  opt.on("-i", "--input FILE", "Input file") { |f| options[:input] = f }
end.parse!

content = File.read(options[:input])
tiles = content.chomp.split("\n")

def dirs(tile)
  tile = tile.split("")
  dirs = []
  while !tile.empty?
    t = tile.shift
    if "ns".include? t
      t += tile.shift
    end
    dirs << t
  end
  dirs
end

def move(point, dir)
  x, y = point
  case dir
  when "e"
    [x+1, y]
  when "w"
    [x-1, y]
  when "nw"
    [x-1, y+1]
  when "ne"
    [x, y+1]
  when "sw"
    [x, y-1]
  when "se"
    [x+1, y-1]
  end
end

def flip(v)
  return :black if v == :white
  :white
end

floor = tiles
  .map { |t| dirs(t) }
  .map { |a| a.inject([0, 0]) { |point, dir| move(point, dir) } }
  .each_with_object(Hash.new(:white)) { |p, m| m[p] = flip(m[p]) }

def count_black(floor)
  floor
    .filter { |k, v| v == :black }
    .count
end

print "P1: "
puts count_black(floor)


def neighbors(p)
  x, y = p
  [
    [x+1, y],
    [x-1, y],
    [x-1, y+1],
    [x, y+1],
    [x, y-1],
    [x+1, y-1],
  ]
end

def iterate(floor)
  neighbor_counts = floor.each_with_object(Hash.new(0)) do |p, nc|
    p, c = p
    if c == :black
      neighbors(p).each { |pn| nc[pn] += 1 }
    end
  end

  neighbor_counts.each_with_object(Hash.new(:white)) do |np, fl|
    p, n = np
    if n == 1 && floor[p] == :black
      fl[p] = :black
    end
    if n == 2
      fl[p] = :black
    end
  end
end

puts "0: #{count_black(floor)}"

for i in (1..100)
  floor = iterate(floor)
  puts "#{i}: #{count_black(floor)}"
end

