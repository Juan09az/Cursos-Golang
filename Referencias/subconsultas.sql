select AVG(c.nota) from cursos c where c.curso='Algebra lineal';

select c.curso, c.nota, e.nombre from cursos c, estudiantes e where c.curso=
    'Algebra lineal' and c.estudiante_id=any(select e.id from estudiantes)
        order by c.nota asc limit 10;
        
mejor promedio: SELECT e.nombre, e.apellido, AVG(c.nota) as promedio FROM cursos c, estudiantes e where c.estudiante_id=
    any(select e.id from estudiantes)  GROUP BY e.nombre, e.apellido order by promedio desc limit 1;

mayores por sexo: select e.nombre, e.apellido, e.edad as ed from estudiantes e
       where e.gender='male' order by e.edad desc;

select e.nombre, e.apellido, e.edad as ed from estudiantes e
       where e.gender='male' order by e.edad desc;


select DISTINCT * from cursos c, estudiantes e where c.estudiante_id=any(select e.id from estudiantes where e.gender='male' ) order by e.edad desc limit 2;

select c.curso, c.nota from cursos c where c.estudiante_id=(select distinct sub1.id from (SELECT  e.id, AVG(c.nota) as promedio FROM cursos c, estudiantes e where c.estudiante_id=any(select e.id from estudiantes) group by e.id order by promedio desc limit 1) as sub1)