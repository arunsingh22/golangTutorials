

🔷 What is PDQSort(Pattern-Defeating Quicksort)?
PDQSort = Quicksort + Introsort + Insertion Sort + Adaptive Tricks

- It’s a hybrid sorting algorithm that:
- Uses quicksort as the core sorting mechanism (for speed),
- Detects bad pivot patterns early (pattern-defeating),
- Falls back to heapsort (like Introsort) when recursion gets too deep,
- Switches to insertion sort for small arrays (typically < 24 elements),
- Adapts to already sorted or partially sorted input.
- It was designed by Orson Peters and is used in:

Rust’s sort_unstable
Go’s sort.Slice

🧠 How PDQSort Defeats Patterns
- PDQSort detects if the input is nearly sorted or has bad pivot behavior (e.g., many repeated elements),
   and Switches to insertion sort if nearly sorted
- Uses randomized pivots or median-of-three to improve pivot quality
- Tracks recursion depth, and falls back to heapsort if needed
- This makes it both fast and robust.

| Algorithm   | Speed (average) | Worst-case   | Stability | Used in                |
| ----------- | --------------- | ------------ | --------- | ---------------------- |
| Quicksort   | ⚡ Very fast     | ❌ O(n²)      | No        | Historically popular   |
| Merge Sort  | ✅ Predictable   | ✅ O(n log n) | ✅ Yes     | Java, Python (Timsort) |
| **PDQSort** | ⚡⚡ Very fast    | ✅ O(n log n) | ❌ No      | Go, Rust               |
| Introsort   | ⚡ Fast          | ✅ O(n log n) | ❌ No      | C++ (`std::sort`)      |
