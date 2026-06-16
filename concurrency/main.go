package concurrency

// Concurency basics.

// func main() {
// unbuferred channel clocks  both, receiver and sender.
// unbuffered chnnels blocks until receive is ready.
// gorutines should be run in seperate functions.
// main is the one gorutine.
// if inside main we send something to the channel we got deadlock. If no one receices.
// Since, both sender and receiver wait until each other, we can prepare in any order, sender or receiver,
// in an unbuferred channel.

// // on this example we prepared sender first
// ch := make(chan string)

// // preapare receiver first. The key, they must be in diff gorutines.
// go func() {
// 	ch <- "hello gorutine"
// }()
// msg := <-ch
// fmt.Println(msg)

// buffered channel does not block the sender until, buffer is full.
// ch := make(chan int, 3)
// ch <- 1
// ch <- 2
// ch <- 3
// //fmt.Println(<-ch, <-ch,  <-ch)

// Range over the channel.
// ch := make(chan int)
// ack := make(chan bool)
// go func() {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 		fmt.Println("sent: ", i)
// 		ack <- true
// 	}
// 	close(ch)
// }()
// for val := range ch {
// 	fmt.Println("received:", val)
// 	<-ack
// }

// Channel direction in functions
// we create simple bidirectional channel and define direction,
// inside gorutine function. not use make() for directioned channel.

// 	ch := make(chan int, 2)
// 	send(ch, 5)
// 	send(ch, 6)
// 	n := receive(ch)
// 	fmt.Println(n)
// }
// func send(in chan<- int, n int) {
// 	in <- n
// }
// func receive(out <-chan int) int {
// 	return <-out
// }

// Context deep-dive.
// Task-1:
/*
Scenario: You have a background worker that generates random numbers indefinitely.
You want your main function to let it run for a moment, then tell it to stop.
ok, We need to generate some time, lets say it is 5 seconds.
option - 1:
	- Create done channel and send signal to this channel from gorutine after 5 seconds.
	- lesten this cahnnel in main, and exit when we got signal from this channel.
option-2:
	- Using context pachage. context.Done() make it for us automatically.
	- create context with time out, or with cancel.
	- run in for loop rangom generator.
	- listen for ctx.Done() with select.
	- when we got done signal exit from loop.
	-

*/
// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	// call worker
// 	// let run 3 seconds.
// 	// then send signal
// 	// wait until go exits.
// 	go worker(ctx)

// 	// lock until ctx done his job.
// 	// The problem here there might be race condition. Main can got done before the gorutine and exit immediately.
// 	// So what can we do?
// 	// maybe wait 1 second after got done? ok lets try this first.
// 	time.Sleep(2 * time.Second)
// 	cancel()
// 	// it works but, need to degug this moment. To avoid gorutine leak.
// 	// to avoid race conditions, between main and worker add waitgroup, and wait in main,  until wokrer is done.
// }
// func worker(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		case <-time.After(200 * time.Millisecond):
// 			fmt.Println("working... ", rand.IntN(10)) // otherwise generate random number.
// 		}
// 	}

// }

// Task-2:
/*
Scenario: Your app calls an external API.
Sometimes the API responds in 50ms, sometimes it takes 2 seconds.
You have a strict 150ms budget.
	- implement fechData() with randomw duration between 50ms..1s.
	- print results. "Success" or "timeout". To see which requiests fail, success.

	Ok, what we can do?
option-1:
	As always, crete done channel and control cancel manually.
	- not efficient.
option-2:
	- create ctx with timeout funciton to control timeouts.
	- create ctx with strict 150ms.
	- run fetch with this ctx. and randomly sumulate job between 50..500 ms.
	- after timeout, when ctx.Done(): timeout message return.
	- if before timeout, send sucess.
	- run 5 times fech() function from main, to simulate api call.

*/

// func main() {
// 	parentCtx := context.Background()

// 	for i := 0; i < 5; i++ {
// 		// we should create fresh timeout for each request.
// 		ctx, cancel := context.WithTimeout(parentCtx, 200*time.Millisecond)
// 		fetch(ctx)
// 		cancel()
// 	}

// 	//<-ctx.Done() -> what is the diff. Wait or not here...
// 	// yes, if we run for loop in gorutine we should wait.. for ctx.Done()
// 	// otherwise main eds before..

// 	/*
// 	 Anyway it is race condition.
// 	 need to add wgroup.
// 	*/

// }
// func fetch(ctx context.Context) {
// 	duration := rand.IntN(450) + 50
// 	select {
// 	case <-ctx.Done():
// 		fmt.Printf("timeout %v: %s\n", duration, ctx.Err().Error())
// 	case <-time.After(time.Duration(duration) * time.Millisecond):
// 		fmt.Printf("Success! -  %vms\n", time.Duration(duration))
// 	}
// }

// ok, now can we do api call in gorutine?
/*
 ok, how we can do this in go?
	wrap for loop to go function.

*/

/*
The Scenario
A user hits your backend API.
This kicks off an Auth Check,
which then kicks off a heavy Database Query.
If the Auth Check fails early, you want to instantly abort the Database Query down the line
so your database doesn't waste precious CPU cycles on an unauthorized request.

Ok, the main idea is cancel parent, so this will cancel all chids.
 - create parent context.
 - create child context from this parent.
 - and if parent is done, close childs too.
*/

// func main() {
// 	parentCtx, parentCancel := context.WithCancel(context.Background())
// 	defer parentCancel()
// 	childCtx, childCancel := context.WithTimeout(parentCtx, 2*time.Second)
// 	defer childCancel()

// 	wg := &sync.WaitGroup{}
// 	wg.Add(2)
// 	//auth check
// 	go func() {
// 		defer wg.Done()
// 		// simulate work.
// 		time.Sleep(150 * time.Millisecond)
// 		// bad password. Fail. Make context done.
// 		fmt.Println("[Auth] Bad password! Cancelling parent context...")
// 		parentCancel()
// 	}()
// 	// long query
// 	go func() {
// 		defer wg.Done()
// 		// its longer than auth check,
// 		// if auth is failing we dont need to run this query. Or break connectoin.
// 		// to do this we must listen the parent context.
// 		duration := 500 * time.Millisecond
// 		select {
// 		case <-childCtx.Done():
// 			fmt.Println("timeout parent:", childCtx.Err().Error())
// 		case <-time.After(duration):
// 			fmt.Println("Success! Affter ", duration)
// 		}
// 	}()

// 	wg.Wait()
// }

// Task-4
/*
 The Scenario
When a user makes a request,
a microservice generates a unique Trace ID (or Request ID).
This ID needs to be attached to every single log line printed by any function processing that request.
Instead of modifying 50 different function signatures to accept a traceID string parameter,
we inject it into the context.Context at the entry point and extract it only where we need to log it.
1. create trace_id in main gorutine.
2. create 3 layer functions layer1->layer2->layer3 sequencally cals one-other.
3. at layer3 extract this "trace_id" value and print as code layer is reached.

*/
// type ctxKey string

// const TRACE_ID ctxKey = "trace_id"

// func main() {
// 	traceID := "trace-999"
// 	ctx := context.WithValue(context.Background(), TRACE_ID, traceID)
// 	layer1(ctx)
// }

// func layer1(ctx context.Context) {
// 	layer2(ctx)
// }

// func layer2(ctx context.Context) {
// 	layer3(ctx)
// }

// func layer3(ctx context.Context) {
// 	if traceId, ok := ctx.Value(TRACE_ID).(string); !ok {
// 		fmt.Println("trace_id key is not found in ctx")
// 	} else {
// 		fmt.Println("I am core layer: trace_id: ", traceId)
// 	}

// }

// ---- Pipeline pattern.
/*
 What is pipeline pattern?
	- It is we combine multiple stages into one component, job.
	- simply we can run each stage sequentally, but each one will wait until prev job is fully done.
	- So, to solve this, I/O pressure etc, we run it in pipeline.
How pipeline logic works?
1. Each pipeline is seperate gorutine.
2. Each take input channel as input.
3. reads from input channel, and do his job.
4. And sends it to output cahnnel,
5. Returns output cahannel.
Main pros - the next stage can start immediately any time curr stage added first value to this out channel.
- Sometimes it myght be owerkill, so use pipleine structure when you:
	- have more stages.
	- you need high throughtput.
	- when you need hande huge amount of data throught stages, without going out of memory.
There are 2 main, thing here, always pass ctx to stages.
and listen ctx Done to avoid memory gorutine leak.
- Because if next stage, consumer is gone already we dont nede to send outpput, and stop.
- Otherwise without ctx it will block forewer sending to curr stage output, when no one reads from this.
- always use receive-only channels for input/output.
- always close output channel in curr stage after sending all needed data.
simple body for stage function is:
-
*/
// func stage(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for data := range in {
// 			// do some staff there and send to out channel.
// 			out <- data
// 		}
// 		close(out) // Always close.
// 	}()
// 	return out
// }
/*
 ok, now lets doo some tasks.
*/

/*
TASK-1:
You are building a basic data-scrubbing pipeline.
You have a raw list of integers, and you want to filter out all odd numbers before passing them to the final consumer.
1. Generate nums.
2. Filter even nums from list.
3. Consume, print filtered nums. (can use main)
*/
// func main() {
// 	nums := []int{}
// 	for range 10 {
// 		nums = append(nums, rand.IntN(50))
// 	}
// 	fmt.Println(nums)
// 	numbers := generate(nums)
// 	filtered := filter(numbers)

// 	for n := range filtered {
// 		fmt.Print(n, " ") // we safe closed will break automatically.
// 	}
// 	// Do there we have reace conditioin?
// 	// We dont because we iterate all numbers until channel is closing.
// 	// When channel is closing? When we done with all nums. so no need to wait.
// }

// func generate(nums []int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out) // ussually best practice.. safe when above logic panics.
// 		for _, n := range nums {
// 			out <- n
// 		}
// 	}()
// 	return out
// }

// func filter(in <-chan int) <-chan int {
// 	out := make(chan int)

// 	go func() {
// 		defer close(out)
// 		for num := range in {
// 			if num%2 != 0 {
// 				out <- num
// 			}
// 		}
// 	}()
// 	return out
// }

/*
TASK-2:
The Scenario
"What if the user cancels the request mid-stream?"
This task requires you to make your pipeline safe against goroutine leaks using Go contexts.
Assume use canceled after geration stage, so we dont need filter anymore but in our previous code we still
filtering.
So we need context whith cancel in real world.
In our taks we just need with some timeout. To simulate cancellation process.
Or if transformer taks is gone we dont need to send to our cahnnel nothing.
*/
/*
 1. Craete context 150ms timeout in main,
 2. run generate, whcih produces words endless, with slow stream..
 3. run transformer, whcih makes uppercase this words.
 4. after timeout both should be stopped. aka must listen context done().
*/
// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	wg := sync.WaitGroup{}

// 	words := []string{"apple", "banana", "orange"}
// 	stream := generate(ctx, words)
// 	transformed := transform(ctx, stream)

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		time.Sleep(150 * time.Millisecond)
// 		cancel()
// 	}()

// 	for t := range transformed {
// 		fmt.Print(t, " ")
// 	}

// 	wg.Wait()

// }

// func generate(ctx context.Context, words []string) <-chan string {
// 	out := make(chan string)
// 	go func() {
// 		defer close(out)
// 		for _, w := range words {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			case <-time.After(100 * time.Millisecond):
// 				out <- w
// 			}
// 		}
// 	}()
// 	return out
// }

// func transform(ctx context.Context, in <-chan string) <-chan string {
// 	out := make(chan string)
// 	go func() {
// 		defer close(out)
// 		for w := range in {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			case out <- strings.ToUpper(w):
// 				// do nothing, just sended succesfully.
// 				// this is best practice.
// 				// if we use default if we cannot sent to out. it will block forever.
// 				// always use on case ..
// 			}
// 		}
// 	}()
// 	return out
// }

/*
Task 3: The Robust Processor (Pipeline Error Propagation)
The Concept
To pass errors down a pipeline safely, senior engineers stop passing raw types (like int or string) through channels.
Instead, they wrap the data and the error together into a single Result struct envelope.
1. Creatae result type. with number, err.
2. Crete generate []result from given nums.
	on specific number set some error.
3. consume from generate and check results for error.
4. if error print error.
*/

// type Result struct {
// 	Num int
// 	Err error
// }

// func main() {
// 	nums := []int{1, 2, 5, 3, 20, 40}
// 	res := generate(nums)

// 	for r := range res {
// 		if r.Err != nil {
// 			fmt.Println(r.Err.Error())
// 			continue
// 		}
// 		fmt.Println(r.Num)
// 	}
// }

// func generate(nums []int) <-chan Result {
// 	out := make(chan Result)
// 	go func() {
// 		defer close(out)
// 		for _, n := range nums {
// 			r := Result{
// 				Num: n,
// 			}
// 			if n == 30 {
// 				r.Err = fmt.Errorf("blacklisted number %d", n)
// 			}
// 			out <- r
// 		}
// 	}()
// 	return out
// }

/*
	Task-4
	Production log parser:
	You are writing a backend utility to process application logs.
	You need to read raw strings, parse them into structured objects, filter for errors, and print them out.
	If a line is corrupted, you cannot silently drop it; you must pass the parsing error down the pipeline so main can track it.

	type LogEntry struct {
	    Timestamp time.Time
	    Level     string
	    Message   string
	}

	type LogResult struct {
	    Entry LogEntry
	    Err   error
	}

	 pipeline order:
		readLines -> parse -> filterError -> Print in (main).

- Use ctx to handle  deadlines, and user cancecellation.
- All errors, should be handled.
*/
// type LogEntry struct {
// 	Timestamp time.Time
// 	Level     string
// 	Message   string
// }
// type LogResult struct {
// 	Log *LogEntry
// 	Err error
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	// send log lines to redline down from pipeline.
// 	logs := []string{
// 		"2026-06-11T10:00:00Z INFO Server started successfully", // Should be filtered out (INFO)
// 		"2026-06-11T10:00:05Z ERROR Database connection failed", // Should be printed (Valid ERROR)
// 		"corrupted raw text line here",                          // Should print Malformed Line Error
// 		"2026-06-11T99:99:99Z ERROR Invalid time layout clock",  // Should print Timestamp Parsing Error
// 		"2026-06-11T10:00:10Z ERROR Write failed: disk full",    // Should be printed (Valid ERROR)
// 	}
// 	lines := readLines(ctx, logs)
// 	parsed := parse(ctx, lines)
// 	filtered := filter(ctx, parsed)

// 	timer := time.AfterFunc(100*time.Millisecond, func() {
// 		fmt.Println("[TIMER] 100ms reached! Injecting context cancellation...")
// 		cancel()
// 	})
// 	defer timer.Stop()

// 	for log := range filtered {
// 		if log.Err != nil {
// 			fmt.Printf("bad line: %v\n", log.Err)
// 			// cancel()
// 			// fmt.Println("Context canceled, must stop all gorutines...")
// 		} else {
// 			fmt.Printf("log: {%v,%s,%s}\n", log.Log.Timestamp, log.Log.Level, log.Log.Message)
// 		}
// 	}

// }

// /*
// 1. Read lines
// 2. Wrap with log result.
// 3. send one-by-one to pipleine/ otuput channel.
// */
// func readLines(ctx context.Context, lines []string) <-chan string {
// 	out := make(chan string)
// 	go func() {
// 		defer close(out)
// 		for _, line := range lines {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			case out <- line:
// 			}
// 			time.Sleep(50 * time.Millisecond)
// 		}
// 	}()
// 	return out
// }

// func parse(ctx context.Context, in <-chan string) <-chan LogResult {
// 	out := make(chan LogResult)
// 	go func() {
// 		defer close(out)
// 		for line := range in {
// 			logResult := LogResult{}
// 			parts := strings.SplitN(line, " ", 3)
// 			if len(parts) < 3 {
// 				logResult.Err = fmt.Errorf("malformed line: '%s'", line)
// 			} else {
// 				ts, err := time.Parse(time.RFC3339, parts[0])
// 				if err != nil {
// 					logResult.Err = fmt.Errorf("parse err: %v", err)
// 				} else {
// 					logResult.Log = &LogEntry{
// 						Timestamp: ts,
// 						Level:     parts[1],
// 						Message:   parts[2],
// 					}
// 				}
// 			}

// 			select {
// 			case <-ctx.Done():
// 				return
// 			case out <- logResult:
// 				// we send on case:, to avoid blocking.
// 			}
// 		}
// 	}()
// 	return out
// }

// // Discard INFO,WARNING
// func filter(ctx context.Context, in <-chan LogResult) <-chan LogResult {
// 	out := make(chan LogResult)
// 	go func() {
// 		defer close(out)
// 		for lr := range in {
// 			if lr.Err == nil && lr.Log.Level != "ERROR" {
// 				continue
// 			}

// 			select {
// 			case <-ctx.Done():
// 				return
// 			case out <- lr:
// 				// avoid blocking.
// 			}
// 		}
// 	}()
// 	return out
// }

/*
Next topic: Fan-out, Fan-in pattern.
Fan out fan in pattern, creates x workers with its own channel, each worker does its work and
sends data back to its assigned channel.
 - Fan-out -> mutiple gorutines with its own channels reads data from one input channel, and
	creates multiple channels with partial data.
 - Fan-in -> collectsa all data from multiple channels and sends it into one output stream channel.
Consumers does not care how mane workers, how many channels here, they just read all data from only
one output stream channel.
Similar to worker pull. But diff.
 - Worker pool:
	- one shared channel between all workers.
	- no need to merge channels stage.
 - Fan-out/Fan-in:
	- seperate channel for each gorutine.
	- need merge channels into one output stream at the end.
Does order of data guranteed?
	- Wroker pools -> not guranteed.
	- Fan-out/Fan-in -> not guranteed.
How should keep order if matters?
	- Wrap data with index, set unique index for each sending data.
	- Sort data by index at the end
*/
/*
 Task-1:
	The Scenario
You have three independent legacy services dumping telemetry numbers into three separate unbuffered
channels.
You need to create a single unified dashboard channel that reads from all of them concurrently.
1. create 3 channels in main.
2. create merge function, which merges channels data, and publishes to out channel.
3. in main read and print from out channel.

*/

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)
// 	ch3 := make(chan string)

// 	go func() {
// 		defer close(ch1)
// 		defer close(ch2)
// 		defer close(ch3)
// 		for i := range 50 {
// 			s := fmt.Sprintf("hi-%d", i)
// 			if i%2 == 0 {
// 				ch2 <- fmt.Sprintf("ch2-%s", s)
// 			} else if i%3 == 0 {
// 				ch3 <- fmt.Sprintf("ch3-%s", s)
// 			} else {
// 				ch1 <- fmt.Sprintf("ch1-%s", s)
// 			}
// 		}
// 	}()

// 	result := merge(ch1, ch2, ch3)

// 	for r := range result {
// 		fmt.Println(r)
// 	}
// }

// func merge(channels ...<-chan string) <-chan string {
// 	out := make(chan string)

// 	wg := &sync.WaitGroup{}
// 	for _, ch := range channels {
// 		wg.Add(1)
// 		go func(c <-chan string) {
// 			defer wg.Done()
// 			for v := range c {
// 				out <- v
// 			}

// 		}(ch)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()
// 	return out
// }

/*
Let's jump straight into Task 2: The Parallel Thumbnail Generator.
This challenge takes what you just learned about merge (Fan-In) and pairs it with
multiple parallel workers reading from the exact same channel (Fan-Out).

 The Architecture Blueprint
- Instead of one stage feeding the next stage directly, we are going to branch out:
- main creates an imageIDProducer channel.
- main invokes worker three separate times, passing that exact same channel to all three.
- Go's runtime will automatically distribute the work evenly among them!
- main collects those 3 worker output channels and feeds them into your bulletproof merge function.

1. create id producer channel,
2. create 3 workers, which is producing to this same channel.
3. each worker returns its own output channel.
4. Main collects all data form thouse output channels.
*/

// func main() {
// 	start := time.Now().UTC()

// 	ids := make([]string, 10)
// 	for i := range ids {
// 		ids[i] = fmt.Sprintf("id-%d", i)
// 	}

// 	idStream := idProducer(ids)

// 	workers := make([]<-chan string, 3)
// 	for i := range workers {
// 		workers[i] = worker(idStream)
// 	}

// 	results := merge(workers)

// 	for r := range results {
// 		fmt.Println(r)
// 	}

// 	fmt.Println("Total time: ", time.Since(start))
// 	fmt.Println("Done!")
// }

// func idProducer(ids []string) <-chan string {
// 	out := make(chan string)
// 	go func() {
// 		defer close(out)
// 		for _, id := range ids {
// 			out <- id
// 		}
// 	}()
// 	return out
// }

// func worker(in <-chan string) <-chan string {
// 	out := make(chan string)
// 	go func() {
// 		defer close(out)
// 		for v := range in {
// 			time.Sleep(100 * time.Millisecond)
// 			out <- v
// 		}
// 	}()
// 	return out
// }

// func merge(workers []<-chan string) <-chan string {
// 	out := make(chan string)
// 	wg := &sync.WaitGroup{}
// 	for _, worker := range workers {
// 		wg.Add(1)
// 		go func(w <-chan string) {
// 			defer wg.Done()
// 			for val := range w {
// 				out <- val
// 			}
// 		}(worker)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()
// 	return out
// }

/*
Task 3 Specification: The Resilient Web Scraper
Your objective is to design a high-throughput, fault-tolerant web scraping pipeline from scratch. This system must utilize the Fan-Out/Fan-In pattern to process URLs concurrently, propagate errors gracefully using a data envelope, and shut down instantly without leaking any background goroutines if a critical failure occurs.

📋 The Architecture & Requirements
1. The Data Envelope (ScrapeResult)
You cannot pass raw string data down the pipeline because errors will occur.

Define a struct named ScrapeResult that contains fields for the target URL (string),
the scraped Data (string), and any encountered Err (error).
1. Create required structs, ScarapeRequest and ScrapeResult
2. Create Request procuder to scrapeReqStream channel.
3. Create X workers all reading from saame stream, and returns its own channels.
4. Merge those channels, and return output channels as result.
5. Print results from out channel.
- The-re 3th website should give error, simulate it.
- After error gorutines must die without leaks.
- use context for to avoid leaking.
*/

// type ScrapeRequest struct {
// 	URL     string
// 	Timeout time.Duration
// }

// type ScrapeResponse struct {
// 	URL   string
// 	Data  *string
// 	Error error
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	requests := []ScrapeRequest{
// 		{URL: "http://google.com", Timeout: 2 * time.Second},
// 		{URL: "http://github.com", Timeout: 2 * time.Second},
// 		{URL: "http://malicious.com", Timeout: 1 * time.Second}, // The failure trigger
// 		{URL: "http://golang.org", Timeout: 2 * time.Second},
// 		{URL: "http://bytelearn.dev", Timeout: 3 * time.Second},
// 	}

// 	reqStream := produce(ctx, requests)

// 	workers := make([]<-chan ScrapeResponse, 3)
// 	for i := range workers {
// 		workers[i] = worker(ctx, reqStream)
// 	}

// 	results := merge(ctx, workers)

// 	for r := range results {
// 		if r.Error != nil {
// 			fmt.Println("Failed!-URL: ", r.URL, " data: ", nil, " error: ", r.Error)
// 			cancel()
// 		} else {
// 			fmt.Println("Success!-URL: ", r.URL, " data: ", *r.Data, "error: ", r.Error)
// 		}
// 	}
// 	fmt.Println("Done!")
// }

// func produce(ctx context.Context, reqs []ScrapeRequest) <-chan ScrapeRequest {
// 	out := make(chan ScrapeRequest)
// 	go func() {
// 		defer close(out)
// 		for _, r := range reqs {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			case out <- r:
// 			}
// 		}
// 	}()
// 	return out
// }

// func worker(ctx context.Context, reqStream <-chan ScrapeRequest) <-chan ScrapeResponse {
// 	out := make(chan ScrapeResponse)
// 	go func() {
// 		defer close(out)
// 		for req := range reqStream {
// 			if ctx.Err() != nil {
// 				return
// 			}
// 			res := ScrapeResponse{
// 				URL: req.URL,
// 			}
// 			time.Sleep(req.Timeout)
// 			if req.URL == "http://malicious.com" {
// 				res.Error = fmt.Errorf("security blocks: malicious site detected")
// 			} else {
// 				s := fmt.Sprintf("Data from: %s", req.URL)
// 				res.Data = &s
// 			}
// 			out <- res
// 		}
// 	}()
// 	return out
// }

// func merge(ctx context.Context, workerChannels []<-chan ScrapeResponse) <-chan ScrapeResponse {
// 	out := make(chan ScrapeResponse)
// 	wg := &sync.WaitGroup{}
// 	for _, wChan := range workerChannels {
// 		wg.Add(1)
// 		go func(ch <-chan ScrapeResponse) {
// 			defer wg.Done()
// 			for val := range ch {
// 				select {
// 				case <-ctx.Done():
// 					return
// 				case out <- val:
// 				}
// 			}
// 		}(wChan)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()
// 	return out
// }
