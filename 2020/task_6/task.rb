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

def solve_second
  questions = {}
  group_size = 0

  total_questions = 0
  File.open('input.txt', 'r').each_line do |line|
    line = line.strip
    if line.empty?
      total_questions += questions.filter { |_question, count| count == group_size }.length

      questions.clear
      group_size = 0
    else
      group_size += 1
      line.each_char do |char|
        questions[char] = questions.fetch(char, 0) + 1
      end
    end
  end
  total_questions + questions.filter { |_question, count| count == group_size }.length
end


if $PROGRAM_NAME == __FILE__
  puts solve_first
  puts solve_second
end
