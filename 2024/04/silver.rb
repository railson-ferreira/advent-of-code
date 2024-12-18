Coordinates = Struct.new(:row, :column)

def main
  raw_input = File.read("input.txt")
  matrix = format_input raw_input
  all_x_coordinates = find_all_x_coordinates matrix
  occurrences_sum = 0
  all_x_coordinates.each do |x_coordinate|
    occurrences_sum += count_xmas_occurrences(matrix, x_coordinate)
  end
  puts occurrences_sum
end

def format_input(raw_input)
  raw_input.split("\n").map do |line|
    line.split("")
  end
end

def find_all_x_coordinates(matrix)
  all_coordinates = []
  matrix.each_with_index do |row, row_index|
    row.each_with_index do |cell, columns_index|
      if cell == "X"
        coordinates = Coordinates.new
        coordinates.row = row_index
        coordinates.column = columns_index
        all_coordinates.push(coordinates)
      end
    end
  end
  all_coordinates
end

def count_xmas_occurrences(matrix, x_coordinate)
  count = 0
  %w[N NE E SE S SW W NW].each do |dir|
    ok = true
    next_coordinates = x_coordinate
    %w[M A S].each do |char|
      next_coordinates = get_next(next_coordinates, dir)
      valid_coordinates = next_coordinates.row >= 0 && next_coordinates.column >= 0
      unless valid_coordinates && matrix.dig(next_coordinates.row, next_coordinates.column) == char
        ok = false
        break
      end
    end
    if ok
      count += 1
    end
  end
  count
end

def get_next(coordinates, dir)
  new_coordinates = Coordinates.new
  case dir
  when "N"
    new_coordinates.row = coordinates.row - 1
    new_coordinates.column = coordinates.column
  when "NE"
    new_coordinates.row = coordinates.row - 1
    new_coordinates.column = coordinates.column + 1
  when "E"
    new_coordinates.row = coordinates.row
    new_coordinates.column = coordinates.column + 1
  when "SE"
    new_coordinates.row = coordinates.row + 1
    new_coordinates.column = coordinates.column + 1
  when "S"
    new_coordinates.row = coordinates.row + 1
    new_coordinates.column = coordinates.column
  when "SW"
    new_coordinates.row = coordinates.row + 1
    new_coordinates.column = coordinates.column - 1
  when "W"
    new_coordinates.row = coordinates.row
    new_coordinates.column = coordinates.column - 1
  when "NW"
    new_coordinates.row = coordinates.row - 1
    new_coordinates.column = coordinates.column - 1
  else
    raise "Invalid direction: #{dir}. Must be one of those: [N, NE, E, SE, S, SW, W, NW]"
  end
  new_coordinates
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"