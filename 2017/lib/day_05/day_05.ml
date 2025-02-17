let input = Utils.read_whole_file "./lib/day_05/input.txt" |> String.split_on_char '\n' 
let total = (List.length input)
let instructions_part1 = Hashtbl.create total
let instructions_part2 = Hashtbl.create total

let() = List.iteri (fun i e -> 
  Hashtbl.add instructions_part1 (i + 1) (int_of_string e);
  Hashtbl.add instructions_part2 (i + 1) (int_of_string e);) input

let rec apply_instructions ins pos step is_part2 = 
  let current_value = Hashtbl.find ins pos in
  let new_pos = pos + current_value in
  let new_value = current_value + if not is_part2 then 1 else if current_value < 3 then 1 else -1 in

  match new_pos with
  | x when x < 1 || x > total -> step
  | _ -> 
    Hashtbl.replace ins pos new_value;
    apply_instructions ins new_pos (step + 1) is_part2

let part1 = string_of_int @@ apply_instructions instructions_part1 1 1 false
let part2 = string_of_int @@ apply_instructions instructions_part2 1 1 true
