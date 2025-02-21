defmodule Day01 do
  def get_numbers() do
    lines = File.stream!("input.txt") |> Stream.map(&String.trim_trailing/1) |> Enum.to_list()
    Enum.map(lines, fn ln -> Integer.parse(ln) |> elem(0) end)
  end

  def part1(numbers) do
    Enum.reduce(numbers, 0, fn n, acc -> n + acc end)
  end

  def part2([head | tail], acc, history) do
    total = acc + head

    if Enum.member?(history, total) do
      total
    else
      part2(tail, total, [total | history])
    end
  end

  def part2([], acc, history) do
    part2(get_numbers(), acc, history)
  end
end

numbers = Day01.get_numbers()
IO.puts("Day 1, part 1 = #{Day01.part1(numbers)}")
IO.puts("Day 1, part 1 = #{Day01.part2(numbers, 0, [])}")
