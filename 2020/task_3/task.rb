def solve_first
  area = []
  File.open('input.txt', 'r').each_line do |line|
    area.push line
  end

  trees = 0
  x = 0
  y = 0
  loop do
    y += 1
    break if area.length == y

    line = area[y]
    x = (x + 3) % (line.length - 1)

    trees = trees.next if line[x] == '#'
  end

  trees
end

puts solve_first
