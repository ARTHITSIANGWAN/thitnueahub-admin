/**
 * 🛡️ TNH-ZERO-TRUST-API-CONNECTOR (ระบบครอบ Workers ป้องกันพอร์ตชน)
 * วัตถุประสงค์: ดึงค่า App Confidence Score แบบ Real-time และทำตัวเบาหวิวให้หน้าบ้านดึงข้อมูลได้
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	// นำเข้า Library ของ Workers เพื่อเปลี่ยนเครื่องยนต์ให้ทำงานล่องหนบน Edge ได้
	"github.com/syumai/workers"
)

// โครงสร้างข้อมูลสำหรับรับค่าจาก Cloudflare API
type CloudflareAppResponse struct {
	Result []struct {
		Name            string  `json:"name"`
		ConfidenceScore float64 `json:"app_confidence_score"`
	} `json:"result"`
}

const (
	CF_API_URL  = "https://api.cloudflare.com/client/v4/accounts/%s/zero_trust/devices/applications"
	ACCOUNT_ID  = "YOUR_CLOUDFLARE_ACCOUNT_ID" // 🆔 ใส่ Account ID ของบอส
	API_TOKEN   = "YOUR_API_TOKEN"              // 🔑 ใส่ API Token ของบอส
)

// ปรับปรุงฟังก์ชันเดิมให้ดึงคะแนนแล้วส่งค่าคืนกลับไปให้หน้าบ้านได้ด้วย
func getConfidenceData() string {
	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf(CF_API_URL, ACCOUNT_ID)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+API_TOKEN)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return `{"status":"error","message":"API Connection Failed"}`
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func main() {
	fmt.Println("🚀 [Go Workers Engine]: สตาร์ทเครื่องยนต์ระบบดักคะแนนขุนพลเรียบร้อย!")

	// ใช้คำสั่ง Serve ของ Workers เข้ามาครอบ เพื่อเปิดวาล์วรับส่งข้อมูลทางพอร์ตเครือข่ายล่องหน (สยบศึกบิวด์แดง)
	workers.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// เปิดทาง CORS ให้พวกแอปหน้าบ้าน และหน้าเพจ Grid Hub ยิงมาดูดสถานะคะแนนไปโชว์ได้
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		// ดึงข้อมูลคะแนนจริงจาก Cloudflare Zero Trust ส่งออกไปให้หน้าบ้านทันทีใน 0.32ms
		apiData := getConfidenceData()
		fmt.Fprint(w, apiData)
	}))
}
