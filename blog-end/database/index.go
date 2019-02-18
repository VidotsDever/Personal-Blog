package database

import (
	"fmt"
	"time"
)

func VerifyCredential(username string, password string) (bool, error) {
	stmt, err := db.Prepare("SELECT password FROM users WHERE username=?")
	if err != nil {
		fmt.Printf("VerifyCredential - %v", err)
		return false, err
	}
	defer stmt.Close()
	var pwd string
	err = stmt.QueryRow(username).Scan(&pwd)
	if err != nil  {
		fmt.Printf("VerifyCredential - %v", err)
		return false, err
	}
	//验证密码是否正确,暂且不加密
	return password == pwd, nil
}

// 传进来的status：draft表示保存，publish表示发布
//　数据库里的status：０代表保存，1代表发布
func SaveOrPublishPost(id, title, htmlStr, rawStr, status string )(bool, error) {
	//　查询是否存在id的文章
	stmt, err := db.Prepare("SELECT * FROM posts WHERE id=?")
	defer stmt.Close()
	if err != nil {
		fmt.Printf("SaveOrPublishPost - %v", err)
		return false, err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Printf("SaveOrPublishPost - %v", err)
		return false, err
	}
	statusCode := 0
	if status == "publish" {
		statusCode = 1
	}
	//表明数据库里没有对应id的文章
	if !rows.Next() {
		stmt,err := db.Prepare("INSERT INTO posts (id, title, html, raw, status, create_time) VALUES (?, ?, ?, ?, ?, ?)")
		if err != nil {
			fmt.Printf("SaveOrPublishPost - %v", err)
			return false, err
		}
		create_time := time.Now().Unix()
		_, err = stmt.Exec(id, title, htmlStr, rawStr, statusCode, create_time)
		if err != nil {
			fmt.Printf("SaveOrPublishPost - %v", err)
			return false, err
		}
		return true, nil
	}
	//更新文章
	if statusCode == 1 {
		stmt, err = db.Prepare("UPDATE posts SET title=?, html=?, raw=?, status=1 WHERE id=?")
	} else {
		stmt, err = db.Prepare("UPDATE posts SET title=?, html=?, raw=? WHERE id=?")
	}

	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(title, htmlStr, rawStr, id)

	if err != nil {
		return false, err
	}

	return true, nil
}

// 获取博客数据：传递的参数 page pagesize
func GetPosts(page, pagesize int) (int,[]Post, error)  {
	stmt, err := db.Prepare("SELECT COUNT(*) FROM posts")
	if err != nil {
		fmt.Println("GetPosts - %v", err)
		return 0, nil, err
	}
	var count int
	err = stmt.QueryRow().Scan(&count)
	if err != nil {
		fmt.Println("GetPosts - %v", err)
		return 0, nil, err
	}
	posts := []Post{}
	//　如果没有博客
	if count == 0 {
		return 0, posts, nil
	}
	offset := (page - 1) * pagesize
	limit := pagesize
	stmt, err = db.Prepare("SELECT id, title, html, raw, create_time FROM posts LIMIT ?, ?")
	if err != nil {
		fmt.Println("GetPosts - %v", err)
		return 0, nil, err
	}
	rows, err := stmt.Query(offset, limit)
	if err != nil {
		fmt.Println("GetPosts - %v", err)
		return 0, nil, err
	}
	var id, title, html, raw string
	var create_time int
	for rows.Next() {
		err = rows.Scan(&id, &title, &html, &raw, &create_time)
		if err != nil {
			fmt.Println("GetPosts - %v", err)
			continue
		}
		post := Post{
			ID: id,
			Title: title,
			HtmlStr: html,
			RawStr: raw,
			CreateTime: create_time,
		}
		posts = append(posts, post)
	}
	return count, posts, nil
}

//　根据id获取文章
func GetPostByID(id string) (Post, error) {
	stmt, err := db.Prepare("SELECT title, html, raw, create_time FROM posts WHERE id=?")
	if err != nil {
		fmt.Println("GetPostByID - %v", err)
	}
	var title, html, raw string
	var create_time int
	row := stmt.QueryRow(id)
	err = row.Scan(&title, &html, &raw, &create_time)
	if err != nil {
		fmt.Println("GetPostByID - %v", err)
	}
	//如果没有数据的话，字段都是默认值
	return Post{
		ID: id,
		Title: title,
		HtmlStr: html,
		RawStr: raw,
		CreateTime: create_time,
	}, nil
}

//　保存用户
func SaveUser(id, avatar, name string) error {
	stmt, err := db.Prepare("SELECT * FROM users WHERE id=?")
	if err != nil {
		fmt.Println("SaveUser %v", err)
		return err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Println("SaveUser %v", err)
		return err
	}
	if rows.Next() {
		//如果用户存在，则直接返回
		fmt.Println("Need not SaveUser")
		return nil
	}
	stmt, err = db.Prepare("INSERT INTO users (id, avatar, name) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("SaveUser %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, avatar, name)
	if err != nil {
		fmt.Println("SaveUser %v", err)
		return err
	}
	return nil
}

// 保存评论
func SaveComment(content, post_id, commenter_id, replyer_id, replyer_name string, parent_id int) (int, error) {
	stmt, err := db.Prepare("INSERT INTO comments (content, post_id, commenter_id, parent_id, replyer_id, replyer_name, create_time) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println("SaveComment---%v", err)
		return 0, err
	}
	defer stmt.Close()
	create_time := time.Now().Unix()
	_, err = stmt.Exec(content, post_id, commenter_id, parent_id, replyer_id, replyer_name, create_time)
	if err != nil {
		fmt.Println("SaveComment---%v", err)
		return 0, err
	}
	stmt, err = db.Prepare("SELECT last_insert_id()")
	if err != nil {
		fmt.Println("SaveComment---%v", err)
		return 0, err
	}
	var id int
	err = stmt.QueryRow().Scan(&id)
	if err != nil {
		fmt.Println("SaveComment---%v", err)
		return 0, err
	}
	return id, nil
}

//　获取评论
func GetComments(post_id string) *[]Comment {
	stmt, err := db.Prepare("SELECT comments.comment_id, comments.commenter_id, comments.parent_id, comments.create_time, comments.replyer_id, comments.replyer_name, comments.content, users.avatar, users.name FROM comments LEFT JOIN users ON comments.post_id=? and comments.commenter_id=users.id")
	if err != nil {
		fmt.Println("GetComments---%v", err)
		return nil
	}
	rows, err := stmt.Query(post_id)
	if err != nil {
		fmt.Println("GetComments---%v", err)
		return nil
	}
	var commenter_id, replyer_id, replyer_name, content, avatar, name string
	var comment_id, parent_id, create_time int
	comments := make([]Comment, 0)
	for rows.Next() {
		err = rows.Scan(&comment_id, &commenter_id, &parent_id, &create_time, &replyer_id, &replyer_name, &content, &avatar, &name)
		if err != nil {
			fmt.Println("Next %v", err)
			continue
		}
		comments = append(comments, Comment{
			Post_ID: post_id,
			Comment_ID: comment_id,
			Commenter_ID: commenter_id,
			Parent_ID: parent_id,
			Create_time: create_time,
			Replyer_id: replyer_id,
			Replyer_name: replyer_name,
			Content: content,
			Commenter_avatar:avatar,
			Commenter_name: name,
			Sub_comments: make([]Comment, 0),
		})
	}

	return  OrderComments(comments)
}

func OrderComments(unordercomments []Comment) *[]Comment {
	comments := make([]Comment, 0)
	for i := 0; i < len(unordercomments); i++  {
		for j := 0; j < len(unordercomments); j++ {
			if unordercomments[i].Comment_ID == unordercomments[j].Parent_ID {
				unordercomments[i].Sub_comments = append(unordercomments[i].Sub_comments, unordercomments[j])
			}
		}
		if unordercomments[i].Parent_ID == 0 {
			comments = append(comments, unordercomments[i])
		}
	}
	return &comments
}