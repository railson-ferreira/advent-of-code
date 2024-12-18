def main
  raw_input = File.read("input.txt")
  instructions_arguments = raw_input.scan(/mul\((\d{1,3}),(\d{1,3})\)|do\(()\)|don't\(\)/)
  enabled = true
  sum = 0
  instructions_arguments.each do |instruction_arguments|
    if instruction_arguments[2] == "" # do
      enabled = true
    end
    if instruction_arguments == [nil, nil, nil] # don't
      enabled = false
    end
    if enabled
      sum += instruction_arguments[0].to_i * instruction_arguments[1].to_i
    end
  end
  puts sum
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"