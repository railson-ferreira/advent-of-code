def main
  raw_input = File.read("input.txt")
  safe_reports_count = 0
  raw_input.split("\n").each do |line|
    parts = line.split(" ")
    report = parts.map{|part|part.to_i}
    if is_safe? report
      safe_reports_count += 1
    else
      report.each_with_index do |_, index|
        new_try = slice(report, index)
        if is_safe? new_try
          safe_reports_count += 1
          break
        end
      end
    end
  end
  puts safe_reports_count
end

def is_safe?(report)
  is_increasing = report[1] > report[0]
  report[1..].each_with_index do |item, index|
    previous = report[index]
    if is_increasing && item < previous
      return false
    end
    if !is_increasing && item > previous
      return false
    end
    if item == previous || diff(item, previous) > 3
      return false
    end
  end
  true
end

def diff(a, b)
  (a - b).abs
end

def slice(array, slice)
  if slice == 0
    return array[1..]
  end
  array[0..slice-1] + array[slice+1..]
end

main