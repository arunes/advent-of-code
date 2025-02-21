defmodule Day02 do
  def get_input() do
    File.stream!("./lib/day-02/input.txt")
    |> Stream.map(&String.trim_trailing/1)
    |> Enum.to_list()
  end

  defp number_of_groups([], twos, threes) do
    {twos, threes}
  end

  defp number_of_groups([head | tail], twos, threes) do
    _chars = head |> String.to_charlist()
    number_of_groups(tail, twos + 1, threes + 2)
  end

  def part1() do
    lines = get_input()
    {twos, threes} = number_of_groups(lines, 0, 0)
    result = twos * threes
    IO.puts(result)
  end
end
