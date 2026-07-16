package main

import (
	"fmt"
	"regexp"
	"strings"
)

// 🧬 Message Filter Interface เพื่อกำหนดขอบเขตหน้าที่ให้ชัดเจน
type MessageFilter interface {
	Clean(input string) string
	IsValid(input string) bool
}

// 🏛️ SecureGate คือ Guard Core หลักของระบบ ThitNueaHub
type SecureGate struct {
	SecretPhrase string
	DangerWords  []string
}

// 🧹 Clean ทำหน้าที่ล้างพวกสัญลักษณ์ Emoji และตัวคั่นปั่นประสาท [🫰][🤔]
func (s *SecureGate) Clean(input string) string {
	// ใช้ RegEx ลบ Emoji และสัญลักษณ์ในวงเล็บเหลี่ยมออกไปให้หมด
	reg := regexp.MustCompile(`\[.*?\]`)
	cleaned := reg.ReplaceAllString(input, "")
	
	// ลบช่องว่างส่วนเกินเพื่อการเปรียบเทียบคำที่แม่นยำ
	return strings.TrimSpace(cleaned)
}

// 🔒 IsValid ทำหน้าที่ตรวจจับคำสั่งแฝงแฝง (เช่น {{stop_broadcast}}) และเช็ครหัสลับเปิดระบบ
func (s *SecureGate) IsValid(input string) bool {
	cleaned := s.Clean(input)

	// 1. ตรวจสอบว่ามีคำสั่งล้างสมอง/คำสั่งหยุดการทำงานจากภายนอกปนมาไหม
	for _, word := range s.DangerWords {
		if strings.Contains(cleaned, word) {
			return false // บล็อกทันทีถ้าเจอคำสั่งอันตราย
		}
	}

	// 2. ตรวจสอบว่ามี "รหัสผ่านสัจจะเปิดระบบ" หรือไม่ (ถ้าไม่มีบอทจะไม่ตอบรับ)
	if !strings.Contains(cleaned, s.SecretPhrase) {
		return false
	}

	return true
}

func main() {
	// ตั้งค่า Core Gatekeeper ป้องกันหลังบ้านสไตล์ ThitNueaHub
	gate := &SecureGate{
		SecretPhrase: "Kin-Khao-Leaw", // รหัสลับสำหรับฝั่งเรา
		DangerWords:  []string{"{{stop_broadcast}}", "{{unsubscribed}}", "{{abort}}"},
	}

	// --- 🧪 การทดสอบระบบคัดกรอง ---

	// เคสที่ 1: ข้อความปั่นประสาทจากบอทภายนอก (หวังให้ระบบเราหยุดทำงาน)
	testInput1 := "สน[🫰]ใจ[🤔]แต่[🥰]บ่สนใจ {{stop_broadcast}}"
	
	// เคสที่ 2: ข้อความส่งคำสั่งจากพวกเราเองอย่างถูกต้อง
	testInput2 := "Kin-Khao-Leaw แอดมินรันระบบ Step 15 ด่วนเลยเน้อ"

	fmt.Println("--- 🛡️ระบบตรวจสอบ ThitNueaHub Active ---")

	// ทดสอบเคสที่ 1
	if gate.IsValid(testInput1) {
		fmt.Printf("เคสที่ 1: ผ่านด่าน (ข้อความจริง: %s)\n", gate.Clean(testInput1))
	} else {
		fmt.Printf("เคสที่ 1: [บล็อก!] ตรวจพบรอยเท้าบอทแฝงเร้นหรือไม่มีรหัสผ่านที่ถูกต้อง\n")
	}

	// ทดสอบเคสที่ 2
	if gate.IsValid(testInput2) {
		fmt.Printf("เคสที่ 2: [ผ่านด่าน] ข้อความสะอาดพร้อมประมวลผล: %s\n", gate.Clean(testInput2))
	} else {
		fmt.Printf("เคสที่ 2: บล็อก!\n")
	}
}
