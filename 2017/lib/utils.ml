let read_whole_file filename =
  let ch = open_in_bin filename in
  let s = really_input_string ch (in_channel_length ch) in
  close_in ch;
  s |> String.trim

let int_value_of_char c = int_of_char c - int_of_char '0'
