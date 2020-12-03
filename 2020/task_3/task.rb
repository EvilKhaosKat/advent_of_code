def read_input
  area = []
  File.open('input.txt', 'r').each_line do |line|
    area.push line
  end

  area
end

def iterate(area, dx, dy)
  trees = 0
  x = 0
  y = 0
  loop do
    y += dy
    break if y >= area.length

    line = area[y]
    x = (x + dx) % (line.length - 1)

    trees = trees.next if line[x] == '#'
  end

  trees
end

def solve_first(area)
  iterate(area, 3, 1)
end

def solve_second(area)
  iterate(area, 1, 1) * iterate(area, 3, 1) * iterate(area, 5, 1) * iterate(area, 7, 1) * iterate(area, 1, 2)
end

puts solve_first read_input
puts solve_second read_input
