lines = File.stream!("input.txt") |> Stream.map(&String.trim_trailing/1) |> Enum.to_list()
total = Enum.reduce(lines, 0, fn ln, acc -> Integer.parse(ln) |> elem(0) |> Kernel.+(acc) end)
IO.puts(total)
