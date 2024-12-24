def main
  raw_input = File.read("input.txt")
  blocks, file_pointers = format_input raw_input
  file_pointers.reverse_each do |file_pointer|
    file_size = get_file_size(blocks, file_pointer)
    empty_pointer = nil
    empty_size = 0
    blocks.each_with_index do |block, index|
      break if index >= file_pointer
      if block.nil?
        empty_pointer = index if empty_pointer.nil?
        empty_size += 1
      else
        empty_pointer = nil
        empty_size = 0
      end
      unless empty_pointer.nil?
        if empty_size >= file_size
          blocks[empty_pointer..empty_pointer - 1 + empty_size] = Array.new(empty_size, blocks[file_pointer])
          blocks[file_pointer..file_pointer - 1 + file_size] = Array.new(file_size, nil)
          # not updating file_pointers[x] since it's not going to be used later
          break
        end
      end
    end
  end
  checksum = 0
  blocks.each_with_index do |block, index|
    next if block.nil?
    checksum += block * index
  end
  puts checksum
end

def format_input(raw_input)
  blocks = []
  file_pointers = []
  raw_input.chars.each_with_index do |char, index|
    is_file = index & 1 == 0
    length = char.to_i
    file_pointers.push blocks.length if is_file
    blocks.concat Array.new(length, is_file ? index / 2 : nil)
  end
  [blocks, file_pointers]
end

def get_file_size(blocks, file_pointer)
  size = 0
  id = blocks[file_pointer]
  blocks[file_pointer..].each do |block|
    break if block != id
    size += 1
  end
  size
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"