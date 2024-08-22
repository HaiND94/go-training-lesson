# **Fan-out**
Fan-out refers to the practice of starting multiple goroutines to handle incoming tasks. The main idea is to distribute incoming tasks to multiple handlers (goroutines) to ensure that each handler deals with a manageable number of tasks.

## Advantages of Fan-out:
- Scalability: By distributing tasks across multiple goroutines, you can efficiently utilize CPU resources and ensure no single goroutine becomes a bottleneck.
- Flexibility: It’s easier to scale up or down by adjusting the number of goroutines as per the workload.

# **Fan-in**
Fan-in is a term used to describe the process of combining multiple results into one channel. By using fan-in, you can simplify how you collect and process data from multiple channels.
## Advantages of Fan-in:
- Simplicity: Centralizing results into a single channel can simplify further processing.
- Efficiency: Managing a single channel can be more efficient than dealing with multiple channels concurrently.

