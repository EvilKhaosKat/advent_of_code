module Part2
  VERSION = "0.1.0"

  def solve_second(input : Array(String))
    #TODO
  end
end

possible_triangles = 0
block = [] of Array(Int32)

File.each_line "../input.txt" do |line| 
  block << line
    .split(' ')
    .reject {|value| value == ""}
    .map {|number| number.to_i }

  if block.size == 3
    (0..2).each do |column_index|
      triangle = [block[0][column_index], block[1][column_index], block[2][column_index]]
      if is_possible_triangle(triangle)
        possible_triangles += 1
      end
    end

    block.clear
  end 
end

def is_possible_triangle(sizes : Array(Int32))
  sizes.sort!
  return (sizes[0] + sizes[1]) > sizes[2]
end

puts possible_triangles