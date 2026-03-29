package dao

import (
	models "blogs/models"
	"log"
)

func GetPostPage(page int, pageSize int) ([]models.Post, error) {
	sql := "SELECT * FROM posts limit ?,?"
	offset := (page - 1) * pageSize
	rows, err := DB.Query(sql, offset, pageSize)
	if err != nil {
		log.Println("查询所有post失败：", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var p models.Post
		err := rows.Scan(
			&p.Pid,
			&p.Title,
			&p.Slug,
			&p.Content,
			&p.Markdown,
			&p.CategoryId,
			&p.UserId,
			&p.ViewCount,
			&p.Type,
			&p.CreateAt,
			&p.UpdateAt,
		)
		if err != nil {
			log.Println("获取所有post失败：", err)
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func GetPostCount() int {
	sql := "select count(*) from posts"
	row := DB.QueryRow(sql)
	var count int
	row.Scan(&count)
	return count
}

func GetPages(total, page, pageSize int) []int {
	if total == 0 || pageSize == 0 {
		return []int{}
	}
	I := total / pageSize
	if total%pageSize != 0 {
		I++
	}
	var pages []int
	for i := 1; i <= I; i++ {
		pages = append(pages, i)
	}
	return pages
}
func GetPostPageByCID(cid, page, pageSize int) ([]models.Post, error) {
	sql := "SELECT * FROM posts where category_id = ? limit ?,? "
	offset := (page - 1) * pageSize
	rows, err := DB.Query(sql, cid, offset, pageSize)
	if err != nil {
		log.Println("查询所有post失败：", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var p models.Post
		err := rows.Scan(
			&p.Pid,
			&p.Title,
			&p.Slug,
			&p.Content,
			&p.Markdown,
			&p.CategoryId,
			&p.UserId,
			&p.ViewCount,
			&p.Type,
			&p.CreateAt,
			&p.UpdateAt,
		)
		if err != nil {
			log.Println("获取所有post失败：", err)
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func GetPostCountByCID(cid int) int {
	sql := "select count(*) from posts where category_id = ?"
	row := DB.QueryRow(sql, cid)
	var count int
	row.Scan(&count)
	return count
}

func GetPostByID(id int) (*models.Post, error) {
	sql := "SELECT * FROM posts WHERE pid = ?"
	row := DB.QueryRow(sql, id)
	if err := row.Err(); err != nil {
		log.Println("获取post失败：", err)
		return nil, err
	}
	p := &models.Post{}
	if err := row.Scan(
		&p.Pid,
		&p.Title,
		&p.Slug,
		&p.Content,
		&p.Markdown,
		&p.CategoryId,
		&p.UserId,
		&p.ViewCount,
		&p.Type,
		&p.CreateAt,
		&p.UpdateAt,
	); err != nil {
		log.Println("获取post失败：", err)
		return nil, err
	}
	return p, nil
}

func SavePost(p *models.Post) error {
	sql := "INSERT INTO posts(title, slug, content, markdown, category_id, user_id, view_count, type, create_at, update_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	ret, err := DB.Exec(sql, p.Title, p.Slug, p.Content, p.Markdown, p.CategoryId, p.UserId, p.ViewCount, p.Type, p.CreateAt, p.UpdateAt)
	if err != nil {
		log.Println("保存post失败：", err)
		return err
	}
	pid, _ := ret.LastInsertId()
	p.Pid = int(pid)
	return nil
}
func UpdatePost(post *models.Post) {
	sql := "UPDATE post SET title = ?, slug = ?, content = ?, markdown = ?, category_id = ?, view_count = ?, type = ?, update_at = ? WHERE pid = ?"
	_, err := DB.Exec(
		sql, post.Title,
		post.Slug, post.Content,
		post.Markdown,
		post.CategoryId,
		post.ViewCount,
		post.Type,
		post.UpdateAt,
		post.Pid,
	)
	if err != nil {
		log.Println("更新文章失败：", err)
		return
	}
}

func GetAllPost() ([]models.Post, error) {
	sql := "SELECT * FROM posts"
	rows, err := DB.Query(sql)
	if err != nil {
		log.Println("获取所有post失败：", err)
		return nil, err
	}
	post := []models.Post{}
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(
			&p.Pid,
			&p.Title,
			&p.Slug,
			&p.Content,
			&p.Markdown,
			&p.CategoryId,
			&p.UserId,
			&p.ViewCount,
			&p.Type,
			&p.CreateAt,
			&p.UpdateAt,
		); err != nil {
			log.Println("获取post失败：", err)
			return nil, err
		}
		post = append(post, p)
	}
	return post, nil
}

func SearchPost(condition string) ([]models.Post, error) {
	sql := `SELECT * FROM posts WHERE title LIKE ? OR content LIKE ?`
	rows, err := DB.Query(sql, "%"+condition+"%", "%"+condition+"%")
	if err != nil {
		log.Println("获取post失败：", err)
		return nil, err
	}
	posts := []models.Post{}
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(
			&p.Pid,
			&p.Title,
			&p.Slug,
			&p.Content,
			&p.Markdown,
			&p.CategoryId,
			&p.UserId,
			&p.ViewCount,
			&p.Type,
			&p.CreateAt,
			&p.UpdateAt,
		); err != nil {
			log.Println("获取post失败：", err)
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}
