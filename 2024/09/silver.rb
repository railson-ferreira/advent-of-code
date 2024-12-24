def main
  raw_input = File.read("input.txt")
  blocks = format_input raw_input
  pointer = blocks.length - 1
  checksum = 0
  blocks.each_with_index do |item, index|
    unless pointer <= index
      if item.nil?
        while blocks[index].nil?
          break if pointer <= index
          blocks[index] = blocks[pointer]
          blocks[pointer] = nil
          pointer -= 1
        end
      end
    end
    break if blocks[index].nil?
    checksum += blocks[index] * index
  end
  puts checksum
end

def format_input(raw_input)
  blocks = []
  raw_input.chars.each_with_index do |char, index|
    is_file = index & 1 == 0
    length = char.to_i
    blocks.concat Array.new(length, is_file ? index / 2 : nil)
  end
  blocks
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"