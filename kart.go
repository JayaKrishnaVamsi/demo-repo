package main

import (
	"fmt"
	"os"
)

/*
type shopping interface {
	Adduser() (bool,[]customer)
	Placeorder()
	Getorder()
	Getproduct()
	Modifyorder()
	Cancelorder()
}
*/
//customer is
type customer struct {
	name   string
	mobile string
	place  string
	orders []Products
}

//Products is...
type Products struct {
	pname string
	qty   int
}

type deliveryrep struct {
	dname  string
	dnum   string
	dplace string
}

//gv
func Adddeliveryrep(n, m, p string, drep []deliveryrep) (bool, []deliveryrep) {
	d := deliveryrep{dname: n, dnum: m, dplace: p}
	for i := 0; i < len(drep); i++ {
		if d.dnum == drep[i].dnum {
			fmt.Println("Number exists")
			return false, drep
		}
	}
	drep = append(drep, d)
	fmt.Println(drep)
	fmt.Println("delivery rep Succesfully registered")

	return true, drep
}

//Adduser is...
func Adduser(n, m, p string, users []customer) (bool, []customer) {
	u := customer{name: n, mobile: m, place: p}
	for i := 0; i < len(users); i++ {
		if u.mobile == users[i].mobile {
			fmt.Println("Number exists")
			return false, users
		}
	}
	users = append(users, u)
	//fmt.Println(users)
	fmt.Println("User Succesfully registered")
	return true, users
}

//Getproduct is
func Getproduct(productlist []Products) {
	fmt.Println("the store offers the following items in the store:")
	fmt.Println("Product", "->", "Quantity")
	for _, j := range productlist {
		fmt.Println(j.pname, "->", j.qty)
	}
}

//Placeorder is
func Placeorder(mob string, pt string, pq int, users []customer, plist []Products) []Products {
	ord := Products{pt, pq}
	for i := 0; i < len(users); i++ {
		if mob == users[i].mobile {
			users[i].orders = append(users[i].orders, ord)
			for j := 0; j < len(plist); j++ {
				if plist[j].pname == pt {
					plist[j].qty -= pq
				}
			}
		}
	}
	fmt.Println(users)
	return plist
}

//Getorder is
func Getorder(mob string, users []customer) {
	for i := 0; i < len(users); i++ {
		if mob == users[i].mobile {
			fmt.Println("Customer ", users[i].name, " ordered the follwing items:")
			fmt.Println(users[i].orders)
		}
	}
}

//Modifyorder is
func Modifyorder(mob string, users []customer) []customer {
	var j int
	for i := 0; i < len(users); i++ {
		if mob == users[i].mobile {
			j = i
			break
		}
	}
	fmt.Println("Modify product type or change?")
	fmt.Println("enter ptype or pqty or both ")
	var t, opt string
	var q int
	fmt.Scanf("%s", &opt)
	fmt.Scanln()
	if opt == "ptype" {
		fmt.Println("enter the new product type")
		fmt.Scanf("%s", &t)
		fmt.Scanln()
		users[j].orders[0].pname = t

	} else if opt == "pqty" {
		fmt.Println("enter the new product Quantity")
		fmt.Scanf("%d", &q)
		fmt.Scanln()
		users[j].orders[0].qty = q
	} else {
		fmt.Println("enter the new product type")
		fmt.Scanf("%s", &t)
		fmt.Scanln()
		users[j].orders[0].pname = t
		fmt.Println("enter the new product Quantity")
		fmt.Scanf("%d", &q)
		fmt.Scanln()
		users[j].orders[0].qty = q
	}
	return users
}

//Cancelorder is
func Cancelorder(mob string, users []customer) []customer {
	var j int
	for i := 0; i < len(users); i++ {
		if mob == users[i].mobile {
			j = i
			break
		}
	}
	fmt.Println("Cancelling the order")
	users[j].orders[0] = users[j].orders[len(users[j].orders)-1] // Copy last element to index i.
	users[j].orders[len(users[j].orders)-1] = Products{}         // Erase last element (write zero value).
	users[j].orders = users[j].orders[:len(users[j].orders)-1]   // Truncate slice.
	return users
}
func main() {
	users := []customer{}
	u1 := customer{name: "vamsi", mobile: "9123456780", place: "a"}
	u2 := customer{name: "krishna", mobile: "1234567890", place: "b"}
	u3 := customer{name: "jaya", mobile: "2345678901", place: "c"}
	users = append(users, u1, u2, u3)
	productlist := []Products{}
	p1 := Products{"tv", 600}
	p2 := Products{"mobile", 800}
	p3 := Products{"laptop", 750}
	p4 := Products{"xbox", 340}
	productlist = append(productlist, p1, p2, p3, p4)
	dreps := []deliveryrep{}
	dr1 := deliveryrep{dname: "lewis", dnum: "2323214527", dplace: "ws"}
	dr2 := deliveryrep{dname: "john", dnum: "2356214527", dplace: "fcv"}
	dr3 := deliveryrep{dname: "carl", dnum: "2323276527", dplace: "ws"}
	dreps = append(dreps, dr1, dr2, dr3)
	var n, m, p string
	var dn, dm, dp string
	fmt.Println("enter name")
	fmt.Scanf("%s", &n)
	fmt.Scanln()
	fmt.Println("enter mobile")
	fmt.Scanf("%s", &m)
	fmt.Scanln()
	fmt.Println("enter place")
	fmt.Scanf("%s", &p)
	var choice int
	for {
		fmt.Println("")
		fmt.Println("---------OPTIONS----------")
		fmt.Println("1.Register")
		fmt.Println("2.Get Product List")
		fmt.Println("3.Place Order")
		fmt.Println("4.Get Orders")
		fmt.Println("5.Modify Order")
		fmt.Println("6.Cancel Order")
		fmt.Println("7.Exit")
		fmt.Println("8.Add Delivery Rep")
		fmt.Println("9.View Delivery Rep")
		fmt.Println("")
		fmt.Println("Enter your choice")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			a, b := (Adduser(n, m, p, users))
			fmt.Println(a)
			users = b
		case 2:
			Getproduct(productlist)
		case 3:
			pt := ""
			pq := 0
			fmt.Println("enter the type you want to buy")
			fmt.Scanf("%s", &pt)
			fmt.Scanln()
			fmt.Println("enter quantity")
			fmt.Scanf("%d", &pq)
			fmt.Scanln()
			p := Placeorder(m, pt, pq, users, productlist)
			productlist = p
			fmt.Println("Remaining Products:", productlist)
		case 4:
			Getorder(m, users)
		case 5:
			users = Modifyorder(m, users)
		case 6:
			users = Cancelorder(m, users)
			fmt.Println("the order has been cancelled", users)
		case 7:
			fmt.Println("Exiting...")
			os.Exit(1)
		case 8:

			fmt.Println("enter drep name")
			fmt.Scanf("%s", &dn)
			fmt.Scanln()
			fmt.Println("enter drep mobile")
			fmt.Scanf("%s", &dm)
			fmt.Scanln()
			fmt.Println("enter drep place")
			fmt.Scanf("%s", &dp)
			a, b := (Adddeliveryrep(dn, dm, dp, dreps))
			fmt.Println(a)
			dreps = b
		case 9:
			fmt.Println(dreps)

		}
	}
}
