# frozen_string_literal: true

# @param [string] line
def get_place(line)
  row_min = 0
  row_max = 127

  column_min = 0
  column_max = 7

  line.each_char.with_index do |char, index|
    if index < 7
      if char == 'F'
        row_max = (row_min + row_max) / 2
      else
        row_min = 1 + (row_min + row_max) / 2
      end
    elsif index < 10
      if char == 'L'
        column_max = (column_min + column_max) / 2
      else
        column_min = 1 + (column_min + column_max) / 2
      end
    end
  end

  [row_min, column_min]
end

def get_id(place)
  place[0] * 8 + place[1]
end

def solve_first
  max_id = 0
  File.open('input.txt', 'r').each_line do |line|
    id = get_id get_place line
    max_id = id if id > max_id
  end

  max_id
end

def solve_second
  ids = []
  File.open('input.txt', 'r').each_line do |line|
    ids.push get_id get_place line
  end

  ids.sort!

  min_id = ids[0]

  min_idx = 0
  max_idx = ids.length - 1

  while min_idx < max_idx
    check_idx = (min_idx + max_idx) / 2

    if ids[check_idx] == check_idx + min_id
      min_idx = 1 + check_idx
    else
      max_idx = check_idx
    end
  end

  ids[min_idx] - 1
end

if $PROGRAM_NAME == __FILE__
  puts solve_first
  puts solve_second
end
