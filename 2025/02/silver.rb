def main
  raw_input = File.read("input.txt")
  ranges = format_input raw_input
  sum_invalid = 0
  ranges.each do |range|
    current, max = range
    while current <= max
      current_str = current.to_s
      if current_str.length.odd?
        current += 1
        next
      end
      half_size = current_str.length / 2
      first_half_current_str = current_str[0..half_size-1]
      last_half_current_str = current_str[half_size..-1]
      sum_invalid += current if first_half_current_str == last_half_current_str
      current += 1
    end
  end
  puts sum_invalid
end

def format_input(raw_input)
  raw_input.split(",").map do |range|
    range.split("-").map(&:to_i)
  end
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"