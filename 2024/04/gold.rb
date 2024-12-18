Coordinates = Struct.new(:row, :column)

def main
  raw_input = File.read("input.txt")
  matrix = format_input raw_input
  all_a_coordinates = find_all_a_coordinates matrix
  occurrences_sum = 0
  all_a_coordinates.each do |a_coordinates|
    if is_x_mas_occurrences(matrix, a_coordinates)
      occurrences_sum += 1
    end
  end
  puts occurrences_sum
end

def format_input(raw_input)
  raw_input.split("\n").map do |line|
    line.split("")
  end
end

def find_all_a_coordinates(matrix)
  all_coordinates = []
  matrix.each_with_index do |row, row_index|
    row.each_with_index do |cell, columns_index|
      if cell == "A"
        coordinates = Coordinates.new
        coordinates.row = row_index
        coordinates.column = columns_index
        all_coordinates.push(coordinates)
      end
    end
  end
  all_coordinates
end

def is_x_mas_occurrences(matrix, a_coordinates)
  %w[NE SE SW NW].each do |dir|
    adjacent_coordinates = get_adjacent(a_coordinates, dir)
    adjacent = matrix.dig(adjacent_coordinates.row, adjacent_coordinates.column)
    if (!is_valid adjacent_coordinates) || adjacent != "M" && adjacent != "S"
      return false
    end
  end
  ne = get_adjacent(a_coordinates, "NE")
  sw = get_adjacent(a_coordinates, "SW")
  if (!is_valid ne) || (!is_valid sw) || matrix.dig(ne.row,ne.column) == matrix.dig(sw.row,sw.column)
    return false
  end
  nw = get_adjacent(a_coordinates, "NW")
  se = get_adjacent(a_coordinates, "SE")
  if (!is_valid nw) || (!is_valid se) || matrix.dig(nw.row,nw.column) == matrix.dig(se.row,se.column)
    return false
  end
  true
end

def get_adjacent(coordinates, dir)
  new_coordinates = Coordinates.new
  case dir
  when "NE"
    new_coordinates.row = coordinates.row - 1
    new_coordinates.column = coordinates.column + 1
  when "SE"
    new_coordinates.row = coordinates.row + 1
    new_coordinates.column = coordinates.column + 1
  when "SW"
    new_coordinates.row = coordinates.row + 1
    new_coordinates.column = coordinates.column - 1
  when "NW"
    new_coordinates.row = coordinates.row - 1
    new_coordinates.column = coordinates.column - 1
  else
    raise "Invalid direction: #{dir}. Must be one of those: [NE, SE, SW, NW]"
  end
  new_coordinates
end

def is_valid(coordinates)
  coordinates.row >= 0 && coordinates.column >= 0
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"