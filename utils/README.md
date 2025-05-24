### TransformRawData2StructpbStruct 使用demo

自定义的slice是无法被解析的 因此需要下面的方法

```go
func (s *BlogService) ListBlog(ctx context.Context, req *pb.ListBlogRequest) (*pb.BlogResponse, error) {
	blogs, pgs, err:=s.uc.GetOnePage(ctx,req.Keyword,int(req.PgNum),int(req.Tid))
	if err!=nil{
		return &pb.BlogResponse{
			Code: 2001,
			Message: "获取博客失败",
		}, nil
	}

	type BlogOutput struct {
		//table各字段
		ID int `json:"id"`
		// Title holds the value of the "title" field.
		Title string `json:"title"`
		// Profile holds the value of the "profile" field.
		// Inset holds the value of the "inset" field.
		Inset string `json:"inset"`
		// View holds the value of the "view" field.
		View int `json:"view"`
		// Up holds the value of the "up" field.
		Up int `json:"up"`
		IsVipContent bool `json:"is_vip_content"`
		IsPublished bool `json:"is_published"`
	}

	var Outputs []*BlogOutput
	for _, v:= range blogs {
		Outputs =append(Outputs, &BlogOutput{
			ID: v.ID,
			Title: v.Title,
			Inset: v.Inset,
			View: v.View,
			Up: v.Up,
			IsVipContent: v.IsVipContent,
			IsPublished: v.IsPublished,
		})
	}

	rawData := map[string]interface{}{
		"real_pgs": int(pgs),
		"blogs":Outputs,
	}

	data, _ := utils.TransformRawData2StructpbStruct(rawData)

	return &pb.BlogResponse{
		Code:    200,
		Message: "success",
		Dt: &pb.BlogResponse_Data{
			Data: structpb.NewStructValue(data),
		},
	}, nil
}
```