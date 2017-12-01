package main

import (
	
	"fmt"
	"math/rand"
	"regexp"
	"time"
	"strings"

)

func ElizaResponse(input string) string{
			
		if matched, _ := regexp.MatchString(`(?!).*\bfather\b.*`,input); matched{
			
			return "Why don’t you tell me more about your father?"
	
		 }
		 re :=  regexp.MustCompile(`(?i)I am([^.?!]*)[.?!]?`)
		 te :=  regexp.MustCompile(`(?i)I'm([^.?!]*)[.?!]?`)
		 ne :=  regexp.MustCompile(`(?i)Im([^.?!]*)[.?!]?`)
		 ye :=  regexp.MustCompile(`(?i)i'm([^.?!]*)[.?!]?`)
		 pe :=  regexp.MustCompile(`(?i)I AM([^.?!]*)[.?!]?`)
		 
		 if matched := re.MatchString(input); matched{
			  
			 return re.ReplaceAllString(input,"How do you know you are $1")
		 
		  }
		 if matched := te.MatchString(input); matched{
			  
			 return te.ReplaceAllString(input,"How do you know you are $1")
		 
		  }
		  if matched := ne.MatchString(input); matched{
			 
			 return ne.ReplaceAllString(input,"How do you know you are $1")
		 
		  }
		  if matched := ye.MatchString(input); matched{
			
			return ne.ReplaceAllString(input,"How do you know that you are $1")
		
		 }
		 if matched := pe.MatchString(input); matched{
			
			return ne.ReplaceAllString(input,"How do you know that you are $1")
		
		 }


         answers := []string{
         
		 "I'm not sure what you're trying to say. Could you explain it to me?",
		 "How does that make you feel?",
		 "Why do you say that?",

		}

		answers2 := []string{

			"How do you know you cannot $1?",
			"I think it perhaps $1 ,if you can try",
			"Can you give a reason why are you think $1 ?",
            "You need believe youself.You can $1.",
		}

		response1 :=  regexp.MustCompile(`(?i)I can't([^.?!]*)[.?!]?`)
		if matched := response1.MatchString(input) ; matched{
		    
			return answers2[rand.Intn(len(answers2))]
			
		}

		response2 :=  regexp.MustCompile(`[Ii] don't(.*)`)
		if matched := response2.MatchString(input) ; matched{
			user := Reflect(input)
			reply := answers2[rand.Intn(len(answers2))]
			return fmt.Sprintf(user, reply)
		}
         

        return answers[rand.Intn(len(answers))]

}


func Reflect(input string) string {

		boundaries := regexp.MustCompile(`\b`)
		tokens := boundaries.Split(input,-1)

		reflections := [][]string {

				{`I`,`you`},
				{`me`,`you`},
				{`you`,`me`},
				{`my`,`your`},
				{`your`,`my`},


		}

		for i,token := range tokens {
				for _,reflection := range reflections {
						if matched,_ := regexp.MatchString(reflection[0], token); matched {

								tokens[i] = reflection[1]
								break
						}
				}
		}

		return strings.Join(tokens,` `)
}





//main function

func main() {

	    rand.Seed(time.Now().UTC().UnixNano()) // Try changing this number!

		
		fmt.Println("I AM exciting.")
		fmt.Println(ElizaResponse("I AM exciting."))
		fmt.Println()
		

		fmt.Println("People say I look like both my mother and father.")
		fmt.Println(ElizaResponse("People say I look like both my mother and father."))
		fmt.Println()

		
		fmt.Println("Father was a teacher.")
		fmt.Println(ElizaResponse("Father was a teacher."))
		fmt.Println()

		fmt.Println("I was my father’s favourite.")
		fmt.Println(ElizaResponse("I was my father’s favourite."))
		fmt.Println()

		fmt.Println("I'm looking forward to the weekend.")
		fmt.Println(ElizaResponse("I'm looking forward to the weekend."))
		fmt.Println()

		fmt.Println("My grandfather was French!")
		fmt.Println(ElizaResponse("My grandfather was French!"))
		fmt.Println()

		fmt.Println("I am happy.")
		fmt.Println(ElizaResponse("I am happy."))
		fmt.Println()

        fmt.Println("I am not happy with your responses.")
		fmt.Println(ElizaResponse("I am not happy with your responses."))
		fmt.Println()

		fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
		fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))
		fmt.Println()

		fmt.Println("I am supposed to just take what you’re saying at face value?")
		fmt.Println(ElizaResponse("I am supposed to just take what you’re saying at face value?"))
		fmt.Println()

		fmt.Println("I am not sure that you understand the effect your questions are having on me.")
		fmt.Println(ElizaResponse("I am not sure that you understand the effect your questions are having on me."))
		fmt.Println()

		fmt.Println("You are my friend.")
		fmt.Println(ElizaResponse("You are my friend."))
		fmt.Println()

		fmt.Println("I don't know how to do my homework.")
		fmt.Println(ElizaResponse("I don't know how to do my homework."))
		fmt.Println()
		
		fmt.Println("I can't believe this news.")
		fmt.Println(ElizaResponse("I can't believe this news."))
		fmt.Println()

}
