def main
  raw_input = File.read("input.txt")
  stones = format_input raw_input
  25.times do
    skip_next = false
    stones.each_with_index do |stone, index|
      if skip_next
        skip_next = false
        next
      end
      if stone == 0
        stones[index] = 1
        next
      end

      stone_str = stone.to_s
      length = stone_str.chars.length
      has_even_number_of_digits = length % 2 == 0
      if has_even_number_of_digits
        half_length = length / 2
        first_half = stone_str[0..half_length - 1].to_i
        last_half = stone_str[half_length..].to_i
        stones[index] = last_half
        stones.insert(index, first_half)
        skip_next = true
      else
        stones[index] *= 2024
      end
    end
  end
  puts stones.length
end

def format_input(raw_input)
  raw_input.split(" ").map do |number|
    number.to_i
  end
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"