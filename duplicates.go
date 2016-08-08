package main

import (
	"fmt"
)

func findDuplicates() {
	fmt.Println("Duplicate keys:")

	// Extract all the unique public keys, with the counts
	rows, err := db.Query(
		`select
            cert_content::jsonb->'FingerprintSHA1' as fingerprint,
            count(cert_content) as count
        from certificates
        where
            cert_content::jsonb->'FingerprintSHA1' IS NOT NULL
        group by
            cert_content::jsonb->'FingerprintSHA1'
        having
            count(cert_content) > 1
        order by
            count desc,
            fingerprint asc`,
	)
	defer rows.Close()
	checkErr(err)

	for rows.Next() {
		var fingerprint []byte
		var count int
		err = rows.Scan(&fingerprint, &count)
		checkErr(err)

		fmt.Printf("%d \t %s\n", count, fingerprint)
	}

}
