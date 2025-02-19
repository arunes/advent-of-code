let read_whole_file filename =
  let ch = open_in_bin filename in
  let s = really_input_string ch (in_channel_length ch) in
  close_in ch;
  s |> String.trim

let int_value_of_char c = int_of_char c - int_of_char '0'

let rec string_of_int_list = function 
  | [] -> ""
  | hd :: tl -> string_of_int hd ^ "," ^ string_of_int_list tl

let rec has_list_in_list lst all idx = match all with
  | [] -> -1
  | hd :: tl ->
      let found = List.for_all (fun a -> a) @@ List.mapi (fun i e -> e = (List.nth lst i)) hd in
      if found then idx else has_list_in_list lst tl (idx + 1)

