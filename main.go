package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	apiKey := "Your_API_KEY"
	url := "https://api.openai.com/v1/chat/completions"

	// ข้อมูลที่ส่งไปยัง API
	data := map[string]interface{}{
		"model": "gpt-4o",
		"messages": []map[string]string{
			{"role": "system", "content": "your name is ส้มลิ้ม"},
			{"role": "user", "content": "คุณชื่ออะไรครับ และ luna คืออะไร btc คืออะไร"},
		},
	}

	// แปลงข้อมูลเป็น JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// ตั้งค่าคำขอ
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// ส่งคำขอและรับคำตอบ
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// ตรวจสอบสถานะการตอบสนอง
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Request failed with status %d: %s\n", resp.StatusCode, string(body))
		return
	}

	// อ่านคำตอบ
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// ตรวจสอบข้อมูลใน result["choices"]
	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		fmt.Println("No choices found in response")
		return
	}

	// ดึงข้อความจาก choices
	message, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	if !ok {
		fmt.Println("No message found in choices")
		return
	}

	content, ok := message["content"].(string)
	if !ok {
		fmt.Println("No content found in message")
		return
	}

	// แสดงคำตอบ
	fmt.Println("Response from ChatGPT:", content)
}


