def main
  raw_input = File.read("input.txt")
  left_numbers = []
  right_numbers = []
  raw_input.split("\n").each do |line|
    parts = line.split("   ")
    left  = parts[0]
    right = parts[1]
    insert_in_order(left_numbers, left.to_i)
    insert_in_order(right_numbers, right.to_i)
  end
  sum_diff = 0
  left_numbers.each_with_index do |item,index|
    sum_diff += diff(item, right_numbers[index])
  end
  puts sum_diff
end

def insert_in_order(array, number)
  array.each_with_index do |item, index|
    if number < item
      array[index,0] = number
      return
    end
  end
  array.push number
end

def diff(a, b)
  (a - b).abs
end

main