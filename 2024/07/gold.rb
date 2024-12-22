def main
  raw_input = File.read("input.txt")
  formated_input = format_input raw_input
  sum = 0
  formated_input.each do |item|
    test_value = item[0]
    numbers = item[1]
    if is_equation_true?(test_value, numbers)
      sum += test_value
    end
  end
  puts sum
end

def format_input(raw_input)
  raw_input.split("\n").map do |line|
    parts = line.split(":")
    test_value = parts[0].to_i
    numbers = parts[1].split(" ").map { |num| num.to_i }
    [test_value, numbers]
  end
end

def is_equation_true?(test_value, numbers)
  permutation = 0
  operations_quantity = numbers.length - 1
  last = 3 ** operations_quantity - 1
  until permutation > last
    operations = permutation.to_ternary_array.reverse
    result = numbers[0]
    operations_quantity.times.each do |index|
      case operations[index]
      when "2"
        result += numbers[index + 1]
      when "1"
        result *= numbers[index + 1]
      else
        # "0" or nil
        result = (result.to_s + numbers[index + 1].to_s).to_i
      end
    end
    return true if result == test_value
    permutation += 1
  end
  false
end

class Integer
  def to_ternary_array
    self.to_s(3).chars
  end
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"