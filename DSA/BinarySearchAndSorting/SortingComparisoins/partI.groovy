



🔄 Why Heapsort is Used as a Fallback in QuickSort (in IntroSort)
💡 What is IntroSort?
Introspective Sort = Quicksort + Heapsort + Insertion Sort
Developed by David Musser (also designed std::sort in C++)
It starts with quicksort for speed, but monitors recursion depth.
If recursion depth exceeds a limit (like 2 * log₂(n)), it switches to heapsort.

✅ Reason for Switching: Avoid QuickSort's Worst Case
| Problem with QuickSort                      | Why Heapsort Solves It        |
| ------------------------------------------- | ----------------------------- |
| Poor pivot → O(n²)                          | Heapsort is always O(n log n) |
| Deep recursion → stack overflow             | Heapsort is non-recursive     |
| Input is nearly sorted or repeated elements | Heapsort doesn’t degrade      |


Even though quicksort is faster on average, its worst-case is unacceptable in some scenarios 
— e.g., when sorting user-controlled input or large datasets.
To avoid this, IntroSort switches to heapsort when it detects danger (too deep recursion = bad pivot choices).

⚠️ So Why Not Just Use Heapsort from the Start?
Because:
Heapsort is slower in practice due to poor cache performance.
Quicksort is faster on most inputs, especially with good pivot strategies.
Heapsort is only a safety net — it guarantees O(n log n) only when needed.

So, it's like:
🚗 Use quicksort for fast driving,
🛑 Switch to heapsort when the road gets rough (deep recursion = potential crash)

📌 Summary
| Sort           | Role in IntroSort              | Why It's Used                                 |
| -------------- | ------------------------------ | --------------------------------------------- |
| QuickSort      | Main sorting algorithm         | Fast, cache-friendly, good on average         |
| Heapsort       | Backup when recursion too deep | Guarantees O(n log n) time, avoids worst-case |
| Insertion Sort | For small subarrays            | Very fast on tiny inputs                      |


🔧 Why Not Merge Sort as the Backup?
- Merge sort is stable and O(n log n), but:
    - It’s not in-place (needs O(n) extra space)
    - You’d lose the memory efficiency of quicksort
- Heapsort keeps the in-place nature, making it a more compatible fallback.

Bottom Line:
Heapsort is the “safety net” in hybrid sorts like IntroSort. It ensures sorting always completes in O(n log n) time, 
even if quicksort hits its worst-case recursion depth — all while staying in-place.