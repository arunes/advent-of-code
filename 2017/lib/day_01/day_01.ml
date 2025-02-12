let input = Utils.read_whole_file "./lib/day_01/input.txt"
let char_list = input |> String.trim |> String.to_seq |> List.of_seq
let length = List.length char_list
let half_way = length / 2

let rec day_01_count list =
  match list with
  | [] -> 0
  | head :: next :: tail ->
      let current_val = Utils.int_value_of_char head in
      let next_val = Utils.int_value_of_char next in
      let value = if current_val = next_val then current_val else 0 in
      value + day_01_count (next :: tail)
  | head :: [] ->
      (*last item*)
      let current_val = Utils.int_value_of_char head in
      let first_char = List.hd char_list in
      let next_val = Utils.int_value_of_char first_char in
      let value = if current_val = next_val then current_val else 0 in
      value
  
let get_next_pos pos = 
  let next_pos = pos + half_way in
  if next_pos >= length then next_pos - length else next_pos

let rec day_02_count pos = 
  if pos >= List.length char_list then 0
  else
    let current_elm = List.nth char_list pos in
    let current_val = Utils.int_value_of_char current_elm in
    let next_elm = List.nth char_list (get_next_pos pos) in
    let next_val = Utils.int_value_of_char next_elm in
    let value = if current_val = next_val then current_val else 0 in
    value + day_02_count (pos + 1)

let part1 = string_of_int (day_01_count char_list)
let part2 = string_of_int (day_02_count 0)
