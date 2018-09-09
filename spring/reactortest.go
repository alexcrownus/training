package main

import "github.com/alexcrownus/presentj"

func main() {
	presentj.MavenTest(`
		package test;
		import reactor.core.publisher.Flux;
		import reactor.test.StepVerifier;
		import org.junit.Test;
		
		public class Sample {
		    // {start OMIT
		    @Test
		    public void fluxTest() {
				Flux<String> flux = Flux.just("foo","bar").log();
				StepVerifier.create(flux)
				.expectNext("foo", "bar")
				.verifyComplete();
		    }
		    // end} OMIT
		}
`, "example/pom.xml")
}
