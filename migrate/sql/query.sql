select * from customers c where city = 'Irvine'

select * from customers c 
left join employees e on concat(e.first_name, ' ', e.last_name) = 'Adam Barr'

select * from products p 
left join order_details od ON od.product_id = p.product_id
left join orders o on o.order_id = od.order_id 
left join customers c on c.customer_id = o.customer_id 
where c.company_name = 'Contonso, Ltd'


select * from orders o 
left join shipping_methods sm on sm.shipping_method_id = o.shipping_method_id 
where sm.shipping_method = 'UPS Ground'

select order_id, (o.freight_charge + o.taxes) as total_cost  from orders o 
order by o.ship_date 