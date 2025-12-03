def main
  raw_input = File.read("input.txt")
  ranges = format_input raw_input
  sum_invalid = 0
  ranges.each do |range|
    current, max = range
    while current <= max
      current_str = current.to_s
      if current_str.length == 1
        current += 1
        next
      end
      (2..current_str.length/2).to_a.push(current_str.length).each do |number|
        next if current_str.length % number != 0
        part_size = current_str.length / number
        first_part = current_str[0, part_size]
        failed = false
        (1...number).each do |i|
          if first_part != current_str[i*part_size, part_size]
            failed = true
            break
          end
        end
        next if failed
        sum_invalid += current
        break
      end
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