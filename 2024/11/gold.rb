CacheKey = Struct.new(:number, :blink_times)
MultiplicationCache = {}

def main
  raw_input = File.read("input.txt")
  stones = format_input raw_input
  sum = 0
  stones.each do |number|
    multiplication = get_multiplication_for_blink_times(number, 75)
    sum += multiplication
  end
  puts sum
end

def format_input(raw_input)
  raw_input.split(" ").map do |number|
    number.to_i
  end
end

def get_multiplication_for_blink_times(number, times)
  return 1 if times == 0
  cached_item = MultiplicationCache[CacheKey.new(number, times)]
  return cached_item if cached_item

  if number == 0
    result = get_multiplication_for_blink_times(1, times - 1)
    MultiplicationCache[CacheKey.new(number, times)] = result
    return result
  end

  number_str = number.to_s
  length = number_str.chars.length
  has_even_number_of_digits = length % 2 == 0
  if has_even_number_of_digits
    half_length = length / 2
    first_half = number_str[0..half_length - 1].to_i
    last_half = number_str[half_length..].to_i
    result_a = get_multiplication_for_blink_times(first_half, times - 1)
    result_b = get_multiplication_for_blink_times(last_half, times - 1)
    result = result_a + result_b
    MultiplicationCache[CacheKey.new(number, times)] = result
    return result
  end
  result = get_multiplication_for_blink_times(number * 2024, times - 1)
  MultiplicationCache[CacheKey.new(number, times)] = result
  result
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"