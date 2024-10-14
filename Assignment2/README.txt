Assignment 2
---
The answers for this section are from the program in JSONstreams.go

To build and run:
  go build JSONstreams.go
  ./JSONstreams

1.1:
  	      Time to serialize , Time to Deserialize
				  ---------------------------------------
   10,000          1.82085ms,   8.346269ms
  100,000        14.091675ms,  72.495739ms
1,000,000       117.840136ms, 633.7812ms


It's important to note that there's a difference between a slice and an array
in Go. An array is a static in length, but a slice points to an underlying
array and can be dynamically resized to accomodate more data. Sometimes the
dynamicism is necessary and worth the additional overhead of allocating a new
array, but if the maximum necessary size is known, it may not be worth it. A
slice stores a header to keep track of the underlying array, the capacity, and
the length.

1.2:
  The time complexity for the algorithm is linear or O(N). There is only a
  single loop being used to create and parse the data, so this will not change
  as the size of the input changes

1.3
  The space complexity is the amount of memory necessary for the algorithm. When
  using an array to store the data, it's the size of a Data structure multiplied
  by the number of records being stored plus the size of the header for being
  stored in a slice. Because the Data structure size and header size will not
  change, it's size complexiy is O(N).

  A Data struct is made up of a 8 byte integer and 16 byte string, so the total
  size used will be 24 bytes per a Data struct multiplied by the total Data
  structs being stored plus the size of the header for the slice which will be
  8 bytes for the pointer to the underlying array, 8 bytes to store the length,
  and 8 bytes to store the capacity for 24 additional bytes. With that, the
  totals are: 10,000*24+24 bytes, 100,000*24+24 bytes, and 1,000,000*24+24 bytes
  for the experiments here.

The answers or this section are from the program in tagsJSON.go

To build and run:
  go build tagsJSON.go
  ./tagsJSON

For 2.1 and 2.2, code is in tagsJSON.go along with functions to randomly select
a value for different fields for the MSDSCourse struct

2.3:
  In this instance only 5 courses need to be stored and there is no need for
  dynamic sizing. Because of this, it makes the most sense to use an Array and
  not slice or map. The array is the most space efficient because it does not
  have the overhead of the slice header and it does not carry the risk of
  accidentally triggering an array reallocation. The slice is potentially better
  from a software engineering standpoint as it's a more flexible data structure.
  The data is of a known structure, so a map does not make sense here either. If
  the data was unstructured, it might make sense to read it into a map and then
  serialize that.
