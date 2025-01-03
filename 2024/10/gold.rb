require "set"
Location = Struct.new(:row, :column)

def main
  raw_input = File.read("input.txt")
  matrix, trailheads = format_input raw_input
  sum = 0
  trailheads.each do |trailhead|
    sum += get_score(trailhead.row, trailhead.column, matrix)
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

def get_score(row, column, matrix)
  score = 0
  get_valid_next_nodes(row, column, matrix).each do |location|
    if matrix[location.row][location.column] == 9
      score += 1
    else
      score += get_score(location.row, location.column, matrix)
    end
  end
  score
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