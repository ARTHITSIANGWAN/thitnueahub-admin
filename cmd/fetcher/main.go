package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func main() {
	ctx := context.Background()
	// เปลี่ยน ProjectID ให้ตรงกับของบอส
	client, err := firestore.NewClient(ctx, "thitnueahub-admin-project-id")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// สมมติว่าเก็บข้อมูลไว้ในคอลเลกชัน "conversations"
	collection := client.Collection("conversations")
	
	// 1. นับจำนวนข้อมูลทั้งหมด
	docs, err := collection.Documents(ctx).GetAll()
	if err != nil {
		log.Fatalf("Failed to count: %v", err)
	}
	fmt.Printf("📊 จำนวนบันทึกทั้งหมดใน SatchaDB: %d รายการ\n", len(docs))

	// 2. ดึงรายการล่าสุด (เรียงตามเวลา)
	query := collection.OrderBy("timestamp", firestore.Desc).Limit(1)
	iter := query.Documents(ctx)
	
	doc, err := iter.Next()
	if err == iterator.Done {
		fmt.Println("⚠️ ยังไม่มีข้อมูลบันทึกในคลัง")
		return
	}
	
	data := doc.Data()
	fmt.Println("📝 ข้อมูลบันทึกล่าสุด:")
	fmt.Printf("   - เวลา: %v\n", data["timestamp"])
	fmt.Printf("   - เนื้อหา: %v\n", data["content"])
}
