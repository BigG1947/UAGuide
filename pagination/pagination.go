package pagination

import "database/sql"

type PaginatingListItems interface {
	GetAllCount(db *sql.DB) (int, error)
	GetItemsForPage(db *sql.DB, page int, limit int) error
	GetItemsForPageByParameters(db *sql.DB, page int, limit int, parameters map[string]string) error
	GetCountForUser(db *sql.DB, id int) (int, error)
	GetItemsForPageForUser(db *sql.DB, page int, id int) error
	GetItemsForPageForUserByParameters(db *sql.DB, page int, limit int, parameters map[string]string, id int) error
}

type Paginator struct {
	HasNext     bool
	HasPrev     bool
	NextPage    int
	PrevPage    int
	TotalItems  int
	TotalPage   int
	CurrentPage int
	LimitItems  int
	Items       PaginatingListItems
	PageList    []int
}

func NewPaginator(db *sql.DB, items PaginatingListItems, page int, limitItems int, filters map[string]string) (*Paginator, error) {
	var p Paginator
	var err error

	p.Items = items
	p.CurrentPage = page
	p.LimitItems = limitItems
	if p.TotalItems, err = p.Items.GetAllCount(db); err != nil {
		return nil, err
	}
	p.TotalPage = p.getTotalPage()
	p.getPageList()
	p.hasNext()
	p.hasPrev()
	if p.HasPrev {
		p.PrevPage = p.CurrentPage - 1
	}
	if p.HasNext {
		p.NextPage = p.CurrentPage + 1
	}
	if err = p.getItems(db, filters); err != nil {
		return nil, err
	}

	return &p, nil
}

func NewPersonalPaginator(db *sql.DB, items PaginatingListItems, page int, limitItems int, filters map[string]string, idUser int) (*Paginator, error) {
	var p Paginator
	var err error

	p.Items = items
	p.CurrentPage = page
	p.LimitItems = limitItems
	if p.TotalItems, err = p.Items.GetAllCount(db); err != nil {
		return nil, err
	}
	p.TotalPage = p.getTotalPage()
	p.getPageList()
	p.hasNext()
	p.hasPrev()
	if err = p.getItemsForUser(db, filters, idUser); err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *Paginator) getTotalPage() int {
	if (p.TotalItems % p.LimitItems) != 0 {
		return 1 + (p.TotalItems / p.LimitItems)
	} else {
		return p.TotalItems / p.LimitItems
	}
}

func (p *Paginator) hasNext() {
	p.HasNext = p.TotalPage > p.CurrentPage
}

func (p *Paginator) hasPrev() {
	p.HasPrev = 1 < p.CurrentPage
}

func (p *Paginator) getPageList() {
	var i int

	if p.TotalPage > 8 {
		for i = 1; i <= p.TotalPage; i++ {
			if i < p.CurrentPage-2 || (i == 1 && p.CurrentPage > 2) || i > p.CurrentPage+2 || (i == p.TotalPage && p.CurrentPage < p.TotalPage-2) {
				continue
			} else {
				p.PageList = append(p.PageList, i)
			}
		}
	} else {
		for i = 1; i <= p.TotalPage; i++ {
			p.PageList = append(p.PageList, i)
		}
	}
}

func (p *Paginator) getItems(db *sql.DB, filters map[string]string) error {
	var err error
	if filters != nil {
		if err = p.Items.GetItemsForPageByParameters(db, p.CurrentPage, p.LimitItems, filters); err != nil {
			return err
		}
	} else {
		if err = p.Items.GetItemsForPage(db, p.CurrentPage, p.LimitItems); err != nil {
			return err
		}
	}
	return nil
}

func (p *Paginator) getItemsForUser(db *sql.DB, filters map[string]string, idUser int) error {
	var err error
	if filters != nil {
		if err = p.Items.GetItemsForPageForUserByParameters(db, p.CurrentPage, p.LimitItems, filters, idUser); err != nil {
			return err
		}
	} else {
		if err = p.Items.GetItemsForPageForUser(db, p.CurrentPage, p.LimitItems); err != nil {
			return err
		}
	}
	return nil
}
