#!/usr/bin/env ruby

require 'optparse'

options = {
  input: "test_input"
}

OptionParser.new do |opts|
  opts.on("-i", "--input FILE", "Input file") { |i| options[:input] = i }
end.parse!

c = File.read(options[:input])

p1, p2 = c.split("\n\n")
p1 = p1.split("\n")[1..].map(&:to_i)
p2 = p2.split("\n")[1..].map(&:to_i)

def play(p1, p2)
  c1 = p1.shift
  c2 = p2.shift
  if c1 > c2
    p1 << c1
    p1 << c2
  else
    p2 << c2
    p2 << c1
  end
end

play(p1, p2) while !p1.empty? && !p2.empty?

d = p1 + p2

puts "P1: #{d.zip(d.length.downto(1)).map { |a, b| a * b }.sum}"

# ----- P2 -----

p1, p2 = c.split("\n\n")
p1 = p1.split("\n")[1..].map(&:to_i)
p2 = p2.split("\n")[1..].map(&:to_i)

def win(deck, card1, card2)
  deck << card1
  deck << card2
end

$memo = {}

def play_recursive(p1, p2, top=false)
  orig_hands = [p1.clone, p2.clone]
  if $memo.include? orig_hands
    return $memo[orig_hands]
  end

  hands = []

  while !p1.empty? && !p2.empty?
    # puts "round: #{p1} #{p2}"
    this_hand = [p1.clone, p2.clone]
    if hands.include? this_hand
      return p1.clone if top
      $memo[orig_hands] = 1
      return 1
    end
    hands << this_hand

    c1 = p1.shift
    c2 = p2.shift

    if c1 > p1.length || c2 > p2.length
      # normal round
      if c1 > c2
        win(p1, c1, c2)
      else
        win(p2, c2, c1)
      end
    else
      if play_recursive(p1.clone[0..c1-1], p2.clone[0..c2-1]) == 1
        win(p1, c1, c2)
      else
        win(p2, c2, c1)
      end
    end
  end

  if p1.empty?
    return p2 if top
    $memo[orig_hands] = 2
    return 2
  else
    return p1 if top
    $memo[orig_hands] = 1
    return 1
  end
end

d = play_recursive(p1, p2, true)

puts "P2: #{d.zip(d.length.downto(1)).map { |a, b| a * b }.sum}"
