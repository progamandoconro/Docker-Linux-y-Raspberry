package main    
    import (
          "fmt"
          "time"
            )
    func main() {
     y := time.Now()
     z := y.Hour()

     if  z < 19 && z > 6 {
           fmt.Print( "こにちわ")

    } else {
            fmt.Print( "こんばんわ")            
    }
        fmt.Print( "お名前は何: ")
        var onamae string    
        fmt.Scanln(&onamae)       
        fmt.Print("どうぞよろし湖お願いします"+ onamae + "さん")

        fmt.Print( "漢字学びましょう, ¡ Vamos a aprender Kanji !" + onamae + "さん")
        fmt.Print("Instrucciones (指示, しじ) para usar este programa (プログラム) y aprender Kanji (漢字)")
        fmt.Print(" Debes acertar todas(すべての) las preguntas (質問) parar ganar (勝利,Shōri)")
        fmt.Print( "¿ Cúal Kanji significa: Kanji ?")
        fmt.Print( "1. 漢字 2. 魚 3. 仕事 4. 女" )
        var q1 int
        var q2 int    
        fmt.Scanln(&q1)  
        if q1 == 1 {
        fmt.Print( "La opción 1. 漢字 es la correcta: 漢字 (かんじ). Mira todos los caracteres kanji que tienes que aprender para ganar" )
        time.Sleep(10000 *time.Millisecond)
        kanji := "一丁七万丈三上下不与丐丑且丕世丗丘丙丞両並丨个中丱串丶丸丹主丼丿乂乃久之乍乎乏乕乖乗乘乙九乞也乢乱乳乾亀亂亅了予争亊事二于云互五井亘亙些亜亞亟亠亡亢交亥亦亨享京亭亮亰亳亶人亻什$
        fmt.Print(kanji)
        fmt.Print( "¿ Cúal Kanji significa compañia/empresa ?")
        fmt.Print( "1. 刀 2. 日本 3. 会社 4. 侍" )   
        fmt.Scanln(&q2)  
        if q2 == 3 {
           for {
                fmt.Print( "漢字, 会社, あなたが勝ちます" )
                time.Sleep(10 * time.Millisecond)
        }
        }else {
        fmt.Print("No has acertado, intenta de nuevo (pista, 会 significa reunión o congregación)")
        fmt.Print( "La opción 3.会社, era la correcta: 会社" )
        }
        }else {
        fmt.Print("No has acertado, intenta de nuevo (pista, 子 (こ, ko) significa niño, los niños aprenden 漢字 rápido")
        fmt.Print( "La opción 1. 漢字 era la correcta: 漢字 (かんじ)" )

        }
        }  

