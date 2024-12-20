require 'set'

Scene = Struct.new(:guard, :obstructions, :ordered_obstructions_by_row, :ordered_obstructions_by_column, :width, :height) do
  def tick
    until can_guard_step_forward?
      guard.turn_right!
    end
    guard.teleport_to_the_side_of_the_obstruction(ordered_obstructions_by_row, ordered_obstructions_by_column)
  end

  def slow_tick
    if can_guard_step_forward?
      guard.take_a_step!
    else
      guard.turn_right!
    end
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

  def can_add_obstruction(obstruction)
    obstruction.row >= 0 &&
      obstruction.column >= 0 &&
      obstruction.row < height &&
      obstruction.column < width
  end

  def add_ordered_obstruction!(obstruction)
    if ordered_obstructions_by_row[obstruction.row] == nil
      ordered_obstructions_by_row[obstruction.row] = []
    end
    if ordered_obstructions_by_column[obstruction.column] == nil
      ordered_obstructions_by_column[obstruction.column] = []
    end
    index = obstructions.find_index { |item| item.row >= obstruction.row && item.column >= obstruction.column } || obstructions.size
    obstructions.insert(index, obstruction)

    by_row_index = ordered_obstructions_by_row[obstruction.row].find_index { |number| number >= obstruction.column } || ordered_obstructions_by_row[obstruction.row].size
    ordered_obstructions_by_row[obstruction.row].insert(by_row_index, obstruction.column)

    by_column_index = ordered_obstructions_by_column[obstruction.column].find_index { |number| number >= obstruction.row } || ordered_obstructions_by_column[obstruction.column].size
    ordered_obstructions_by_column[obstruction.column].insert(by_column_index, obstruction.row)
  end

  def remove_obstruction!(obstruction)
    obstructions.delete(obstruction)
    ordered_obstructions_by_row[obstruction.row].delete(obstruction.column)
    ordered_obstructions_by_column[obstruction.column].delete(obstruction.row)
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

  def teleport_to_the_side_of_the_obstruction(ordered_obstructions_by_row, ordered_obstructions_by_column)
    location = nil
    case direction
    when "N"
      ordered_obstructions_by_column[self.location.column]&.reverse_each do |obstruction|
        next if obstruction > self.location.row
        location = Location.new(obstruction + 1, self.location.column)
        break
      end
    when "E"
      ordered_obstructions_by_row[self.location.row]&.each do |obstruction|
        next if obstruction < self.location.column
        location = Location.new(self.location.row, obstruction - 1)
        break
      end
    when "S"
      ordered_obstructions_by_column[self.location.column]&.each do |obstruction|
        next if obstruction < self.location.row
        location = Location.new(obstruction - 1, self.location.column)
        break
      end
    when "W"
      ordered_obstructions_by_row[self.location.row]&.reverse_each do |obstruction|
        next if obstruction > self.location.column
        location = Location.new(self.location.row, obstruction + 1)
        break
      end
    else
      raise "Invalid direction: #{dir}. Must be one of those: [N, E, S, W]"
    end
    if location == nil
      self.location = get_next_location
    else
      self.location = location
    end
  end
end

def main
  raw_input = File.read("input.txt")
  scene = format_input raw_input
  initial_guard = scene.guard.dup
  possible_positions = Set.new
  until scene.is_guard_off_the_map?
    if scene.can_guard_step_forward?
      new_obstruction = scene.guard.get_next_location
      if new_obstruction != initial_guard.location
        if scene.can_add_obstruction new_obstruction
          backup_guard = scene.guard.dup
          scene.guard = initial_guard.dup
          scene.add_ordered_obstruction! new_obstruction
          path = []
          until scene.is_guard_off_the_map?
            path_node = { direction: scene.guard.direction, location: scene.guard.location }
            if path.include?(path_node)
              possible_positions.add new_obstruction
              break
            else
              path.push(path_node)
            end
            scene.tick
          end
          scene.guard = backup_guard
          scene.remove_obstruction! new_obstruction
        end
      end
    end
    scene.slow_tick
  end
  puts possible_positions.length
end

def format_input(raw_input)
  scene = Scene.new
  scene.width = raw_input.split("\n")[0].chars.length
  scene.height = raw_input.split("\n").length
  scene.obstructions = []
  raw_input.split("\n").each_with_index do |line, row_index|
    line.chars.each_with_index do |char, column_index|
      if char == "#"
        obstruction = Location.new
        obstruction.row = row_index
        obstruction.column = column_index
        scene.obstructions.push obstruction
      end
      if char == "^"
        raise "Cannot set guard twice" if scene.guard != nil
        scene.guard = Guard.new
        scene.guard.direction = "N"
        scene.guard.location = Location.new
        scene.guard.location.row = row_index
        scene.guard.location.column = column_index
      end
    end
  end
  scene.ordered_obstructions_by_row = {}
  scene.ordered_obstructions_by_column = {}
  scene.obstructions.each do |obstruction|
    if scene.ordered_obstructions_by_row[obstruction.row] == nil
      scene.ordered_obstructions_by_row[obstruction.row] = []
    end
    if scene.ordered_obstructions_by_column[obstruction.column] == nil
      scene.ordered_obstructions_by_column[obstruction.column] = []
    end
    scene.ordered_obstructions_by_row[obstruction.row].push obstruction.column
    scene.ordered_obstructions_by_column[obstruction.column].push obstruction.row
  end
  scene
end

start_time = Time.now
main
puts "time elapsed: #{(Time.now - start_time) * 1000} ms"