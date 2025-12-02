def main
  raw_input = File.read("input.txt")
  rotations = format_input raw_input
  dial_position = 50
  count_zero_positions = 0
  rotations.each do |rotation|
    direction, amount = rotation
    if direction == "L"
      dial_position -= amount
      while dial_position < 0 do
        dial_position += 100
      end
    else
      dial_position += amount
      while dial_position > 99 do
        dial_position -= 100
      end
    end
    if dial_position == 0
      count_zero_positions += 1
    end
  end
  puts count_zero_positions
end

def format_input(raw_input)
  raw_input.split("\n").map do |line|
    first_char = line[0]
    rest_chars = line[1..-1]
    [first_char, rest_chars.to_i]
  end
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"