def main
  raw_input = File.read("input.txt")
  rotations = format_input raw_input
  dial_position = 50
  count_zero_positions = 0
  rotations.each do |rotation|
    direction, amount = rotation
    started_at_zero = dial_position == 0
    if direction == "L"
      dial_position -= amount
      first_iteration = true
      while dial_position < 0 do
        count_zero_positions += 1 if dial_position != 0 && (!started_at_zero || !first_iteration)
        dial_position += 100
        first_iteration = false
      end
      count_zero_positions += 1 if dial_position == 0
    else
      dial_position += amount
      while dial_position > 99 do
        count_zero_positions += 1
        dial_position -= 100
      end
    end
    # puts "#{direction}#{amount} #{dial_position} #{count_zero_positions}"
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