require "set"
Location = Struct.new(:row, :column)

def main
  raw_input = File.read("input.txt")
  matrix, trailheads = format_input raw_input
  sum = 0
  trailheads.each do |trailhead|
    positions = get_nine_height_positions(trailhead.row, trailhead.column, matrix)
    sum += positions.length
  end
  puts sum
end

def format_input(raw_input)
  trailheads = []
  matrix = raw_input.split("\n").each_with_index.map do |line, row_index|
    line.chars.each_with_index.map do |char, column_index|
      value = char.to_i
      if value == 0
        trailheads.push(Location.new(row_index, column_index))
      end
      value
    end
  end
  [matrix, trailheads]
end

def get_nine_height_positions(row, column, matrix)
  nine_height_positions = Set.new
  get_valid_next_nodes(row, column, matrix).each do |location|
    if matrix[location.row][location.column] == 9
      nine_height_positions.add(location)
    else
      get_nine_height_positions(location.row, location.column, matrix).each do |position|
        nine_height_positions.add(position)
      end
    end
  end
  nine_height_positions
end

def get_valid_next_nodes(row, column, matrix)
  value = matrix[row][column]
  raise "value out of range" if value >= 9
  width = matrix[row].length
  height = matrix.length
  valid_next_nodes = []
  [
    Location.new(row - 1, column),
    Location.new(row, column + 1),
    Location.new(row + 1, column),
    Location.new(row, column - 1)
  ].each do |location|
    next if location.row < 0 || location.column < 0 ||
      location.row >= width || location.column >= height
    if matrix[location.row][location.column] - value == 1
      valid_next_nodes.push(location)
    end
  end
  valid_next_nodes
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"