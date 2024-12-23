require "set"
Location = Struct.new(:row, :column)

def main
  raw_input = File.read("input.txt")
  antennas_by_char, height, width = format_input raw_input
  antinodes_location_set = Set.new
  antennas_by_char.each do |key_value|
    _, locations = key_value
    locations.each_with_index do |location_a, index_a|
      locations[index_a + 1..].each do |location_b|
        antinodes_location_set.add(location_a)
        antinodes_location_set.add(location_b)

        diff = Location.new(location_b.row - location_a.row, location_b.column - location_a.column)
        antinode_a = Location.new(location_a.row - diff.row, location_a.column - diff.column)
        antinode_b = Location.new(location_b.row + diff.row, location_b.column + diff.column)

        while antinode_a.row >= 0 && antinode_a.column >= 0 && antinode_a.row < height && antinode_a.column < width
          antinodes_location_set.add(antinode_a)
          antinode_a = Location.new(antinode_a.row - diff.row, antinode_a.column - diff.column)
        end
        while antinode_b.row >= 0 && antinode_b.column >= 0 && antinode_b.row < height && antinode_b.column < width
          antinodes_location_set.add(antinode_b)
          antinode_b = Location.new(antinode_b.row + diff.row, antinode_b.column + diff.column)
        end
      end
    end
  end
  puts antinodes_location_set.length
end

def format_input(raw_input)
  lines = raw_input.split("\n")
  locations_by_char = {}
  lines.each_with_index do |line, row|
    line.chars.each_with_index do |char, column|
      unless char == "."
        locations_by_char[char] = [] if locations_by_char[char].nil?
        locations_by_char[char].push Location.new(row, column)
      end
    end
  end
  [locations_by_char, lines.length, lines[0].chars.length]
end

def is_equation_true?(test_value, numbers)
  permutation = 0
  operations_quantity = numbers.length - 1
  last = 2 ** operations_quantity - 1
  until permutation > last
    operations = permutation.to_binary_bool_array.reverse
    result = numbers[0]
    operations_quantity.times.each do |index|
      if operations[index] # sum
        result += numbers[index + 1]
      else
        # times
        result *= numbers[index + 1]
      end
    end
    return true if result == test_value
    permutation += 1
  end
  false
end

class Integer
  def to_binary_bool_array
    self.to_s(2).chars.map { |char| char == '1' }
  end
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"