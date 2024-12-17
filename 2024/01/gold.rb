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
  similarity_sum = 0
  left_numbers.each do |item|
    similarity_sum += get_similarity_score(item, right_numbers)
  end
  puts similarity_sum
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

def get_similarity_score(number, ordered_list)
  times = 0
  ordered_list.each do |item|
    next if item < number
    if item > number
      break
    end
    times += 1
  end
  number * times
end

main