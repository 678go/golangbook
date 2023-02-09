# 字符串

1. > [字符串常用操作](craftdocs://open?blockId=99583473-570F-4B9F-B822-C72411DD2DE0&spaceId=650c2b8b-9755-d6cc-4b6f-6e8738fcd089)
    1. > [Search (contains, prefix/suffix, index)](craftdocs://open?blockId=B1B3E911-C49F-4D3C-AB0A-EC6F364581ED&spaceId=650c2b8b-9755-d6cc-4b6f-6e8738fcd089)
    2. > [Replace (uppercase/lowercase, trim)](craftdocs://open?blockId=0A6BE58B-A25A-4EFC-BE36-61B843B9CD9D&spaceId=650c2b8b-9755-d6cc-4b6f-6e8738fcd089)
    3. > [Split by space or comma](craftdocs://open?blockId=64E9B141-B9B1-4D5E-AC98-C1F91A6256E8&spaceId=650c2b8b-9755-d6cc-4b6f-6e8738fcd089)
    4. > [Join strings with separator](craftdocs://open?blockId=133A5F70-C1EC-4BF8-B8AD-8C69719B61CA&spaceId=650c2b8b-9755-d6cc-4b6f-6e8738fcd089)
1. ## 字符串常用操作

### 1.1 Search (contains, prefix/suffix, index)

| **Expression**                         | **Result** | **Note**                         |
| -------------------------------------- | ---------- | -------------------------------- |
| `strings.Contains("Japan", "abc")`     | false      | Is abc in Japan?                 |
| `strings.ContainsAny("Japan", "abc")`  | true       | Is a, b or c in Japan?           |
| `strings.Count("Banana", "ana")`       | 1          | Non-overlapping instances of ana |
| `strings.HasPrefix("Japan", "Ja")`     | true       | Does Japan start with Ja?        |
| `strings.HasSuffix("Japan", "pan")`    | true       | Does Japan end with pan?         |
| `strings.Index("Japan", "abc")`        | \-1        | Index of first abc               |
| `strings.IndexAny("Japan", "abc")`     | 1          | a, b or c                        |
| `strings.LastIndex("Japan", "abc")`    | \-1        | Index of last abc                |
| `strings.LastIndexAny("Japan", "abc")` | 3          | a, b or c                        |

### 1.2 Replace (uppercase/lowercase, trim)

| **Expression**                                                              | **Result** | **Note**                                             |
| --------------------------------------------------------------------------- | ---------- | ---------------------------------------------------- |
| `strings.Replace("foo", "o", ".", 2)`                                       | f..        | Replace first two “o” with “.” Use -1 to replace all |
| `f := func(r rune) rune {  \     return r + 1  \ }  \ strings.Map(f, "ab")` | bc         | Apply function to each character                     |
| `strings.ToUpper("Japan")`                                                  | JAPAN      | Uppercase                                            |
| `strings.ToLower("Japan")`                                                  | japan      | Lowercase                                            |
| `strings.Title("ja pan")`                                                   | Ja Pan     | Initial letters to uppercase                         |
| `strings.TrimSpace(" foo ")`                                                | foo        | Strip leading and trailing white space               |
| `strings.Trim("foo", "fo")`                                                 |            | Strip *leading and trailing* f:s and o:s             |
| `strings.TrimLeft("foo", "f")`                                              | oo         | *only leading*                                       |
| `strings.TrimRight("foo", "o")`                                             | f          | *only trailing*                                      |
| `strings.TrimPrefix("foo", "fo")`                                           | o          |                                                      |
| `strings.TrimSuffix("foo", "o")`                                            | fo         |                                                      |

### 1.3 Split by space or comma

| **Expression**                   | **Result**   | **Note**           |
| -------------------------------- | ------------ | ------------------ |
| `strings.Fields(" a  b ")`       | `["a" "b"]`  | Remove white space |
| `strings.Split("a,b", ",")`      | `["a" "b"]`  | Remove separator   |
| `strings.SplitAfter("a,b", ",")` | `["a," "b"]` | Keep separator     |

### 1.4 Join strings with separator

| **Expression**                          | **Result** | **Note**         |
| --------------------------------------- | ---------- | ---------------- |
| `strings.Join([]string{"a", "b"}, ":")` | a:b        | Add separator    |
| `strings.Repeat("da", 2)`               | dada       | 2 copies of “da” |

### 1.5 Format and convert

| **Expression**               | **Result** | **Note**      |
| ---------------------------- | ---------- | ------------- |
| `strconv.Itoa(-42)`          | `"-42"`    | Int to string |
| `strconv.FormatInt(255, 16)` | `"ff"`     | Base 16       |

2. ## string、int、int64相互转换

