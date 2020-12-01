# frozen_string_literal: true

require 'set'

def solve_first
  numbers = Set.new

  File.open('input.txt', 'r').each_line do |line|
    number = line.to_i
    target_number = 2020 - number

    return number * target_number if numbers.include? target_number

    numbers.add number
  end
end


# TODO:
def solve_second
  numbers = []

  File.open('input.txt', 'r').each_line do |line|
    numbers.push line.to_i
  end
  numbers.sort.reverse!

  for idx_1 in 0...numbers.size
    for idx_2 in idx_1...numbers.size
      for idx_3 in idx_2...numbers.size
        number_1 = numbers[idx_1]
        number_2 = numbers[idx_2]
        number_3 = numbers[idx_3]

        break if number_1 + number_2 > 2020

        return number_1 * number_2 * number_3 if number_1 + number_2 + number_3 == 2020
      end
    end
  end
end

puts solve_first
puts solve_second
