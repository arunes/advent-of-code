let input = Utils.read_whole_file "./lib/day_01/input.txt"
let char_list = input |> String.to_seq |> List.of_seq

let rec count list =
  match list with
  | [] -> 0
  | head :: next :: tail ->
      let current_val = Utils.int_value_of_char head in
      let next_val = Utils.int_value_of_char next in
      let value = if current_val = next_val then current_val else 0 in
      value + count (next :: tail)
  | head :: [] ->
      (*last item*)
      let current_val = Utils.int_value_of_char head in
      let first_char = List.hd char_list in
      let next_val = Utils.int_value_of_char first_char in
      let value = if current_val = next_val then current_val else 0 in
      value

let part1 = string_of_int (count char_list)
