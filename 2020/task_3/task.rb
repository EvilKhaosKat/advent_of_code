def read_input
  area = []
  File.open('input.txt', 'r').each_line do |line|
    area.push line
  end

  area
end

def iterate(area, delta_x, delta_y)
  trees = x = y = 0
  loop do
    y += delta_y
    break if y >= area.length

    line = area[y]
    x = (x + delta_x) % (line.length - 1)

    trees = trees.next if line[x] == '#'
  end

  trees
end

def solve_first(area)
  iterate(area, 3, 1)
end

def solve_second(area)
  iterate(area, 1, 1) \
    * iterate(area, 3, 1) \
    * iterate(area, 5, 1) \
    * iterate(area, 7, 1) \
    * iterate(area, 1, 2)
end

puts solve_first read_input
puts solve_second read_input
