package services

import (
	"fmt"
	"net/url"
	"strings"
)


func VerifyLink(inputURL string)(string,error){
	//verificação básica
	validUrl , err := isValid(inputURL)
	if err != nil || validUrl == ""{
		return ""  , err
	}
	// Segurança
	if !isSafeURL(validUrl){return "" , fmt.Errorf("host não é válido")}
	if hasXSSAttempt(validUrl){return ""  , fmt.Errorf("tentativa de XSS detectada")}

	return validUrl , nil
}


/// 


func urlLength(inputURL string)error{
	if len(inputURL) > 2048{
		return fmt.Errorf("url length ERROR")
	}
	if len(inputURL)< 5{
		return fmt.Errorf("url length ERROR")
	}
	return nil
}

func isSafeURL(inputURL string)bool{
	parsedURL, err := url.Parse(inputURL)
    if err != nil {
        return false
    }

	blockedHosts := []string{
		"localhost" , "127.0.0.1" , "0.0.0.0" , "::1" , "192.168." , "10." , "172.16." , "169.254.",
	}

	hostname := parsedURL.Hostname()

	for _,blocked := range blockedHosts{
		if strings.Contains(hostname , blocked){
			return false
		}
	}

	return true
}

func hasXSSAttempt(inputURL string)bool{
	xssPatterns := []string{
        `javascript:`,
        `data:text/html`,
        `vbscript:`,
        `onload=`,
        `onerror=`,
        `<script`,
        `</script>`,
    }
	lowerURL := strings.ToLower(inputURL)
	for _,pattern := range xssPatterns{
		if strings.Contains(lowerURL , pattern){
			return true
		}
	}

	return false
}

func isValid(inputURL string)(string , error){
	// verifica se possuí http ou https, se não houver adiciona
	if !strings.HasPrefix(inputURL , "http://") && !strings.HasPrefix(inputURL , "https://"){
		inputURL = "https://" + inputURL	
	}

	// tenta fazer o parse da url
	parsedURL , err := url.Parse(inputURL)
	if err != nil{return "", err}

	// verifica o protocolo
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https"{
		return "" , fmt.Errorf("protocolo deve ser http ou https")
	}

	//verifica o host, se vazio, retorna o erro
	if parsedURL.Host == "" {
        return "", fmt.Errorf("host não pode estar vazio")
    }
	// tamanho
	err = urlLength(inputURL)
	if err != nil{
		return "" , err
	}
	return inputURL , nil

}