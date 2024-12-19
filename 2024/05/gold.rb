FormatedInput = Struct.new(:rules, :updates)
Rule = Struct.new(:before, :after)

def main
  raw_input = File.read("input.txt")
  formated_input = format_input raw_input
  rules = formated_input.rules
  updates = formated_input.updates
  middles_sum = 0
  updates.each do |update|
    unless is_correct(rules, update)

      middles_sum += get_middle get_corrected_update(rules, update)
    end
  end
  puts middles_sum
end

def format_input(raw_input)
  formated_input = FormatedInput.new
  formated_input.rules = []
  formated_input.updates = []
  parts = raw_input.split("\n\n")
  rules = parts[0]
  updates = parts[1]
  rules.split("\n").each do |item|
    numbers = item.split("|")
    formated_input.rules.push Rule.new(numbers[0].to_i, numbers[1].to_i)
  end
  updates.split("\n").each do |item|
    update = item.split(",").map do |number|
      number.to_i
    end
    formated_input.updates.push update
  end
  formated_input
end

def is_correct(rules, update)
  filtered_rules = rules.select do |rule|
    update.include?(rule.before) && update.include?(rule.after)
  end
  filtered_rules.each do |rule|
    update.each do |number|
      if number == rule.after
        return false
      end
      if number == rule.before
        break
      end
    end
  end
  true
end

def get_corrected_update(rules, update)
  update = update.dup
  filtered_rules = rules.select do |rule|
    update.include?(rule.before) && update.include?(rule.after)
  end
  until is_correct(rules, update)
    filtered_rules.each do |rule|
      after_index = nil
      update.each_with_index do |number, index|
        if number == rule.after
          after_index = index
        end
        if number == rule.before
          if after_index.nil?
            break
          end
          update.move!(index, after_index)
        end
      end
    end
  end
  update
end

def get_middle(update)
  raise "Array length must be odd!" if update.length.even?
  middle_index = update.length / 2
  update[middle_index]
end

class Array
  def move!(old_index, new_index)
    item = self.delete_at(old_index)
    self.insert(new_index, item)
  end
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"