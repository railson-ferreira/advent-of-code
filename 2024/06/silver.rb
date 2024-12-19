require 'set'

Scene = Struct.new(:guard, :obstructions, :width, :height) do
  def tick
    until can_guard_step_forward?
      guard.turn_right!
    end
    guard.take_a_step!
  end

  def can_guard_step_forward?
    !obstructions.include? guard.get_next_location
  end

  def is_guard_off_the_map?
    guard.location.row < 0 ||
      guard.location.column < 0 ||
      guard.location.row >= height ||
      guard.location.column >= width
  end
end
Location = Struct.new(:row, :column)
Guard = Struct.new(:location, :direction) do
  def turn_right!
    case direction
    when "N"
      self.direction = "E"
    when "E"
      self.direction = "S"
    when "S"
      self.direction = "W"
    when "W"
      self.direction = "N"
    else
      raise "Invalid direction: #{dir}. Must be one of those: [N, E, S, W]"
    end
  end

  def get_next_location
    location = Location.new
    case direction
    when "N"
      location.row = self.location.row - 1
      location.column = self.location.column
    when "E"
      location.row = self.location.row
      location.column = self.location.column + 1
    when "S"
      location.row = self.location.row + 1
      location.column = self.location.column
    when "W"
      location.row = self.location.row
      location.column = self.location.column - 1
    else
      raise "Invalid direction: #{dir}. Must be one of those: [N, E, S, W]"
    end
    location
  end

  def take_a_step!
    self.location = get_next_location
  end
end

def main
  raw_input = File.read("input.txt")
  scene = format_input raw_input
  distinct_positions = Set.new
  until scene.is_guard_off_the_map?
    distinct_positions.add scene.guard.location
    scene.tick
  end
  puts distinct_positions.length
end

def format_input(raw_input)
  scene = Scene.new
  scene.width = raw_input.split("\n")[0].chars.length
  scene.height = raw_input.split("\n").length
  guard = nil
  obstructions = []
  raw_input.split("\n").each_with_index do |line, row_index|
    line.chars.each_with_index do |char, column_index|
      if char == "#"
        obstruction = Location.new
        obstruction.row = row_index
        obstruction.column = column_index
        obstructions.push obstruction
      end
      if char == "^"
        raise "Cannot set guard twice" if guard != nil
        guard = Guard.new
        guard.direction = "N"
        guard.location = Location.new
        guard.location.row = row_index
        guard.location.column = column_index
      end
    end
  end
  scene.guard = guard
  scene.obstructions = obstructions
  scene
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"