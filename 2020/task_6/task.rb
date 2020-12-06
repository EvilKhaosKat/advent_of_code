# frozen_string_literal: true
require 'set'

def solve_first
  questions = Set.new

  total_questions = 0
  File.open('input.txt', 'r').each_line do |line|
    line = line.strip
    if line.empty?
      total_questions += questions.length
      questions.clear
    else
      line.each_char do |char|
        questions.add char
      end
    end
  end
  total_questions + questions.length
end


if $PROGRAM_NAME == __FILE__
  puts solve_first
end
